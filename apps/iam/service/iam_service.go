package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"infra-kit/apps/iam/ent"
	"infra-kit/apps/iam/ent/authority"
	"infra-kit/apps/iam/ent/credential"
	"infra-kit/apps/iam/ent/group"
	"infra-kit/apps/iam/ent/namespace"
	"infra-kit/apps/iam/ent/org"
	"infra-kit/apps/iam/ent/user"
	"infra-kit/lib/redislib"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateUserParam struct {
	OrgId string
	Name  string
	Email string
	sk    string
}

type CreateNamespaceParam struct {
	OrgId string
	Name  string
	Code  string
}

type ListResult struct {
	PageSize      int
	PageTotal     int
	ItemsTotal    int
	NextPageToken string
	Items         any
}

func newListResult(pageSize int, items any) *ListResult {
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 1000
	}
	return &ListResult{
		PageSize: pageSize,
		Items:    items,
	}
}

func (res *ListResult) SetItemsTotal(total int) {
	res.ItemsTotal = total
	if total == 0 {
		res.PageTotal = 0
	} else {
		pageTotal := res.ItemsTotal / res.PageSize
		y := res.ItemsTotal % res.PageSize
		if y > 0 {
			pageTotal += 1
		}
		res.PageTotal = pageTotal
	}
}

func (res *ListResult) SetItems(items any) {
	res.Items = items
}

const (
	KeyAccessToken = "Access-Token"
)

var (
	errInvalidPageToken = errors.New("page token is invalid")
)

type IAMService interface {
	Auth(ctx context.Context, email, pwd string) (token string, expireTime time.Time, err error)
	CheckAuthStatus(ctx context.Context, token string) (orgCode, userId string, err error)
	ValidateAuthorities(ctx context.Context, userId string, authorityIds []string) (map[string]bool, error)

	CreateAuthority(ctx context.Context, code, name string) (*ent.Authority, error)
	CreateGroup(ctx context.Context, orgId, code, name, description string) (*ent.Group, error)
	CreateNamespace(ctx context.Context, param CreateNamespaceParam) (*ent.Namespace, error)
	CreateOrg(ctx context.Context, code, name string) (*ent.Org, error)
	CreateUser(ctx context.Context, param CreateUserParam) (*ent.User, error)
	GetAuthority(ctx context.Context, authorityId string, isOrgCode bool) (*ent.Authority, error)
	GetGroup(ctx context.Context, groupId string, withAuthorityIds bool) (*ent.Group, error)
	GetNamespace(ctx context.Context, id string) (*ent.Namespace, error)
	GetOrg(ctx context.Context, id string) (*ent.Org, error)
	GetUser(ctx context.Context, id string) (*ent.User, error)
	ListAuthority(ctx context.Context, pageSize int, pageToken *string, groupId *string, isOrgCode bool) (*ListResult, error)
	ListGroup(ctx context.Context, pageSize int, pageToken *string, orgId *string) (*ListResult, error)
	ListNamespace(ctx context.Context, pageSize int, pageToken *string, orgId *string) (*ListResult, error)
	ListOrg(ctx context.Context, pageSize int, pageToken *string) (*ListResult, error)
	ListUser(ctx context.Context, pageSize int, pageToken *string, orgId, groupId, nsId *string) (*ListResult, error)
	UpdateCredential(ctx context.Context, userId string, ak, sk *string) (*ent.Credential, error)
	UpdateAuthority(ctx context.Context, id string, code, name *string) (*ent.Authority, error)
	UpdateGroup(ctx context.Context, groupId string, code, name *string, addUserIds, removeUserIds, addAuthorityIds, removeAuthorityIds []string) (*ent.Group, error)
	UpdateNamespace(ctx context.Context, id string, code, name *string) (*ent.Namespace, error)
	UpdateOrg(ctx context.Context, id string, code, name *string) (*ent.Org, error)
	UpdateUser(ctx context.Context, id string, email, name *string) (*ent.User, error)
}

var _ IAMService = (*iamService)(nil)

type iamService struct {
	db    *ent.Client
	cache *redislib.Redis
}

// GetUser implements IAMService.
func (s *iamService) GetUser(ctx context.Context, id string) (*ent.User, error) {
	return s.db.User.Query().Where(user.ID(id)).Only(ctx)
}

// UpdateAuthority implements IAMService.
func (s *iamService) UpdateAuthority(ctx context.Context, id string, code *string, name *string) (*ent.Authority, error) {
	m := s.db.Authority.UpdateOneID(id)
	if code != nil {
		m.SetCode(*code)
	}
	if name != nil {
		m.SetName(*name)
	}
	m.SetMtime(time.Now())
	res, err := m.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateNamespace implements IAMService.
func (s *iamService) UpdateNamespace(ctx context.Context, id string, code *string, name *string) (*ent.Namespace, error) {
	m := s.db.Namespace.UpdateOneID(id)
	if code != nil {
		m.SetCode(*code)
	}
	if name != nil {
		m.SetName(*name)
	}
	m.SetMtime(time.Now())
	res, err := m.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateOrg implements IAMService.
func (s *iamService) UpdateOrg(ctx context.Context, id string, code *string, name *string) (*ent.Org, error) {
	m := s.db.Org.UpdateOneID(id)
	if code != nil {
		m.SetCode(*code)
	}
	if name != nil {
		m.SetName(*name)
	}
	m.SetMtime(time.Now())
	res, err := m.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUser implements IAMService.
func (s *iamService) UpdateUser(ctx context.Context, id string, email *string, name *string) (*ent.User, error) {
	m := s.db.User.UpdateOneID(id)
	if email != nil {
		m.SetEmail(*email)
	}
	if name != nil {
		m.SetName(*name)
	}
	m.SetMtime(time.Now())
	res, err := m.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetNamespace implements IAMService.
func (s *iamService) GetNamespace(ctx context.Context, id string) (*ent.Namespace, error) {
	q := s.db.Namespace.Query().Where(namespace.ID(id))
	return q.Only(ctx)
}

// ListNamespace implements IAMService.
func (s *iamService) ListNamespace(ctx context.Context, pageSize int, pageToken *string, orgId *string) (*ListResult, error) {
	res := newListResult(pageSize, ent.Users{})
	listQuery := s.db.Namespace.Query().
		Order(ent.Desc(namespace.FieldID)).
		Limit(res.PageSize + 1)
	if orgId != nil {
		listQuery.Where(namespace.HasOrgWith(org.ID(*orgId)))
	}
	total, err := listQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	res.SetItemsTotal(total)

	if pageToken != nil {
		bytes, err := base64.StdEncoding.DecodeString(*pageToken)
		if err != nil {
			return nil, errInvalidPageToken
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(namespace.IDLTE(pageToken))
	}

	items, err := listQuery.All(ctx)
	if len(items) == res.PageSize+1 {
		res.NextPageToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%v", items[len(items)-1].ID)))
		items = items[:len(items)-1]
	}
	res.SetItems(items)
	return res, err
}

// ListOrg implements IAMService.
func (s *iamService) ListOrg(ctx context.Context, pageSize int, pageToken *string) (*ListResult, error) {
	res := newListResult(pageSize, ent.Orgs{})
	listQuery := s.db.Org.Query().
		Order(ent.Desc(org.FieldID)).
		Limit(res.PageSize + 1)

	total, err := listQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	res.SetItemsTotal(total)

	if pageToken != nil {
		bytes, err := base64.StdEncoding.DecodeString(*pageToken)
		if err != nil {
			return nil, errInvalidPageToken
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(org.IDLTE(pageToken))
	}

	items, err := listQuery.All(ctx)
	if len(items) == res.PageSize+1 {
		res.NextPageToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%v", items[len(items)-1].ID)))
		items = items[:len(items)-1]
	}
	res.SetItems(items)
	return res, err
}

func (s *iamService) GetGroup(ctx context.Context, groupId string, withAuthorityIds bool) (*ent.Group, error) {
	q := s.db.Group.Query().Where(group.ID(groupId))
	if withAuthorityIds {
		q.WithAuthorities()
	}
	return q.Only(ctx)
}

func (s *iamService) ListGroup(ctx context.Context, pageSize int, pageToken *string, orgId *string) (*ListResult, error) {
	res := newListResult(pageSize, ent.Users{})
	listQuery := s.db.Group.Query().
		Order(ent.Desc(user.FieldID)).
		Limit(res.PageSize + 1)
	if orgId != nil {
		listQuery.Where(group.HasOrgWith(org.ID(*orgId)))
	}
	total, err := listQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	res.SetItemsTotal(total)

	if pageToken != nil {
		bytes, err := base64.StdEncoding.DecodeString(*pageToken)
		if err != nil {
			return nil, errInvalidPageToken
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(group.IDLTE(pageToken))
	}

	items, err := listQuery.All(ctx)
	if len(items) == res.PageSize+1 {
		res.NextPageToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%v", items[len(items)-1].ID)))
		items = items[:len(items)-1]
	}
	res.SetItems(items)
	return res, err
}

// ListAuthority implements IAMService.
func (s *iamService) ListAuthority(ctx context.Context, pageSize int, pageToken, groupId *string, isOrgCode bool) (*ListResult, error) {
	res := newListResult(pageSize, ent.Authorities{})
	listQuery := s.db.Debug().Authority.Query().
		Order(ent.Desc(user.FieldID)).
		Limit(res.PageSize + 1)
	if isOrgCode {
		listQuery.Where(authority.CodeContains("{org_code}"))
	}
	if groupId != nil {
		if *groupId != "" {
			listQuery.Where(authority.HasGroupsWith(group.ID(*groupId)))
		}
	}
	total, err := listQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	res.SetItemsTotal(total)

	if pageToken != nil {
		bytes, err := base64.StdEncoding.DecodeString(*pageToken)
		if err != nil {
			return nil, errInvalidPageToken
		}
		pageToken := string(bytes)
		listQuery = listQuery.Where(authority.IDLTE(pageToken))
	}

	items, err := listQuery.All(ctx)
	if len(items) == res.PageSize+1 {
		res.NextPageToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%v", items[len(items)-1].ID)))
		items = items[:len(items)-1]
	}
	res.SetItems(items)
	return res, err
}

func (s *iamService) GetAuthority(ctx context.Context, authorityId string, isOrgCode bool) (*ent.Authority, error) {
	q := s.db.Authority.Query()
	if isOrgCode {
		q.Where(authority.CodeContains("{org_code}"))
	}
	return q.Where(authority.ID(authorityId)).Only(ctx)
}

func (s *iamService) GetOrg(ctx context.Context, id string) (*ent.Org, error) {
	return s.db.Org.Query().Where(org.ID(id)).Only(ctx)
}

func (s *iamService) UpdateGroup(ctx context.Context, groupId string, code, name *string, addUserIds, removeUserIds, addAuthorityIds, removeAuthorityIds []string) (*ent.Group, error) {
	m := s.db.Group.UpdateOneID(groupId)
	if code != nil {
		m.SetCode(*code)
	}
	if name != nil {
		m.SetName(*name)
	}
	m.AddUserIDs(addUserIds...)
	m.RemoveUserIDs(removeUserIds...)
	m.AddAuthorityIDs(addAuthorityIds...)
	m.RemoveAuthorityIDs(removeAuthorityIds...)
	m.SetMtime(time.Now())
	res, err := m.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *iamService) ListUser(ctx context.Context, pageSize int, pageToken *string, orgId, groupId, nsId *string) (*ListResult, error) {
	res := newListResult(pageSize, ent.Users{})
	listQuery := s.db.User.Query().
		Order(ent.Desc(user.FieldID)).
		Limit(res.PageSize + 1)
	if orgId != nil {
		listQuery.Where(user.HasOrgWith(org.ID(*orgId)))
	}
	if groupId != nil {
		listQuery.Where(user.HasGroupsWith(group.ID(*groupId)))
	} else if nsId != nil {
		listQuery.Where(user.HasNamespaceWith(namespace.ID(*nsId)))
	}
	total, err := listQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	res.SetItemsTotal(total)

	if pageToken != nil {
		bytes, err := base64.StdEncoding.DecodeString(*pageToken)
		if err != nil {
			return nil, errInvalidPageToken
		}
		pageToken := string(bytes)
		listQuery = listQuery.
			Where(user.IDLTE(pageToken))
	}

	items, err := listQuery.All(ctx)
	if len(items) == res.PageSize+1 {
		res.NextPageToken = base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%v", items[len(items)-1].ID)))
		items = items[:len(items)-1]
	}
	res.SetItems(items)
	return res, err
}

// CreateAuthority implements AuthService.
func (s *iamService) CreateAuthority(ctx context.Context, code string, name string) (*ent.Authority, error) {
	now := time.Now()
	res, err := s.db.Authority.Create().
		SetCode(code).
		SetName(name).
		SetCtime(now).
		SetMtime(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *iamService) UpdateCredential(ctx context.Context, userId string, ak, sk *string) (*ent.Credential, error) {
	now := time.Now()
	user, err := s.db.User.Query().Where(user.ID(userId)).WithCredential().Only(ctx)
	if err != nil {
		return nil, err
	}
	if user.Edges.Credential == nil {
		return nil, errors.New("missing credential")
	}
	c := s.db.Credential.UpdateOneID(user.Edges.Credential.ID)

	if ak != nil {
		c.SetAk(*ak)
	}
	if sk != nil {
		c.SetSk(*sk)
	}
	c.SetMtime(now)
	return c.Save(ctx)
}

// CreateGroup implements AuthService.
func (s *iamService) CreateGroup(ctx context.Context, orgId, code, name, description string) (*ent.Group, error) {
	now := time.Now()
	res, err := s.db.Group.Create().
		SetOrgID(orgId).
		SetCode(code).
		SetName(name).
		SetDescription(description).
		SetCtime(now).
		SetMtime(now).
		SetDeleted(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewIAMService(db *ent.Client, cache *redislib.Redis) IAMService {
	return &iamService{
		db:    db,
		cache: cache,
	}
}

// Auth implements UserService.
func (s *iamService) Auth(ctx context.Context, email, pwd string) (token string, expireTime time.Time, err error) {
	now := time.Now()
	q := s.db.User.Query().WithCredential(func(cq *ent.CredentialQuery) { cq.Where(credential.Sk(pwd)) })
	q.Where(user.Email(email))

	q.WithOrg().Only(ctx)
	account, err := q.WithOrg().Only(ctx)
	if ent.IsNotFound(err) {
		return "", time.Time{}, errors.Wrap(err, "user is not found")
	} else if err != nil {
		return
	}

	token = uuid.NewString()
	expire := 30 * time.Minute

	if err = s.cache.Client.Set(ctx, s.getAccessTokenKey(token), fmt.Sprintf("%s:%s", account.Edges.Org.Code, account.ID), expire).Err(); err != nil {
		return
	}

	return token, now.Add(expire), nil
}

func (s *iamService) CheckAuthStatus(ctx context.Context, token string) (orgCode string, userId string, err error) {
	tokenKey := s.getAccessTokenKey(token)
	getter := s.cache.Get(ctx, tokenKey)
	err = getter.Err()
	if err == redis.Nil {
		err = errors.New("access token expired")
		return
	} else if err != nil {
		return
	}

	vals := strings.Split(getter.Val(), ":")
	if len(vals) < 2 {
		err = errors.New("access token is invalid")
		return
	}
	return vals[0], vals[1], nil
}

func (s *iamService) ValidateAuthorities(ctx context.Context, token string, codes []string) (map[string]bool, error) {
	tokenKey := s.getAccessTokenKey(token)
	ttler := s.cache.TTL(ctx, tokenKey)
	ttl := ttler.Val()
	res := make(map[string]bool)
	if len(codes) == 0 {
		return res, nil
	}
	getter := s.cache.Get(ctx, tokenKey)
	err := getter.Err()
	if err == redis.Nil {
		return nil, errors.New("access token expired")
	} else if err != nil {
		return nil, err
	}

	vals := strings.Split(getter.Val(), ":")
	if len(vals) < 2 {
		return nil, errors.New("access token is invalid")
	}

	authoritiesKey := s.getAuthoritiesKey(vals[0], vals[1])
	missingCodes := make([]string, 0)
	for _, code := range codes {
		authorityKey := authoritiesKey + ":" + code
		val, err := s.cache.Get(ctx, authorityKey).Result()
		if err != nil {
			fmt.Printf("get key %s failed:%v", authorityKey, err)
			missingCodes = append(missingCodes, code)

		} else {
			if val == "0" {
				res[code] = false
			} else if val == "1" {
				res[code] = true
			}
		}
	}

	var exists []struct {
		Id string `json:"id,omitempty"`
	}
	authorityIds := make([]string, 0)
	if len(missingCodes) > 0 {
		if err := s.db.Authority.Query().Where(authority.CodeIn(missingCodes...)).Select(authority.FieldID).Scan(ctx, &exists); err != nil {
			return nil, err
		}
	}
	for _, e := range exists {
		authorityIds = append(authorityIds, e.Id)
	}

	authorities, err := s.db.User.Query().Where(user.IDEQ(vals[1])).QueryGroups().WithAuthorities(func(aq *ent.AuthorityQuery) { aq.Where(authority.IDIn(authorityIds...)) }).QueryAuthorities().All(ctx)
	if err != nil {
		return nil, err
	}

	existDict := make(map[string]struct{})
	for _, a := range authorities {
		existDict[a.Code] = struct{}{}
	}

	for _, e := range codes {
		if _, ok := existDict[e]; ok {
			res[e] = true
		}
	}
	userId := vals[1]

	havingAuthorities, err := s.db.User.Query().Where(user.IDEQ(userId)).QueryGroups().WithAuthorities(func(aq *ent.AuthorityQuery) { aq.Where(authority.IDIn(authorityIds...)) }).QueryAuthorities().All(ctx)
	if err != nil {
		return nil, err
	}

	selectIds := make(map[string]struct{}, len(havingAuthorities))
	for _, e := range havingAuthorities {
		selectIds[e.ID] = struct{}{}
	}

	for _, id := range authorityIds {
		if _, ok := selectIds[id]; ok {
			res[id] = true
		} else {
			res[id] = false
		}
	}

	for _, code := range missingCodes {
		val := "0"
		if res[code] {
			val = "1"
		}
		authorityKey := authoritiesKey + ":" + code
		err := s.cache.Set(ctx, authorityKey, val, ttl).Err()
		if err != nil {
			log.Println(err)
		}
	}

	return res, nil
}

// CreateNamespace implements UserService.
func (s *iamService) CreateNamespace(ctx context.Context, param CreateNamespaceParam) (*ent.Namespace, error) {
	now := time.Now()
	res, err := s.db.Namespace.Create().
		SetCode(param.Code).
		SetName(param.Name).
		SetCtime(now).
		SetMtime(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrg implements UserService.
func (s *iamService) CreateOrg(ctx context.Context, code, name string) (*ent.Org, error) {
	now := time.Now()
	res, err := s.db.Org.Create().
		SetCode(code).
		SetName(name).
		SetCtime(now).
		SetMtime(now).
		SetDeleted(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *iamService) CreateUser(ctx context.Context, param CreateUserParam) (*ent.User, error) {
	now := time.Now()
	res, err := s.db.User.Create().
		SetOrgID(param.OrgId).
		SetName(param.Name).
		SetEmail(param.Email).
		SetCredential(&ent.Credential{
			Ak:    "",
			Sk:    param.sk,
			Ctime: now,
			Mtime: now,
		}).
		SetCtime(now).
		SetMtime(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *iamService) getAccessTokenKey(token string) string {
	return fmt.Sprintf("%s:%s", KeyAccessToken, token)
}

func (s *iamService) getAuthoritiesKey(orgCode, userId string) string {
	return fmt.Sprintf("Authorities:%s:%s", orgCode, userId)
}
