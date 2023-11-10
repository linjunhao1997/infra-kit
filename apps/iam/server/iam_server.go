package server

import (
	"context"
	"infra-kit/apps/iam/ent"
	"infra-kit/apps/iam/pb"
	"infra-kit/apps/iam/service"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAMServiceServer struct {
	pb.UnimplementedIAMServiceServer
	service service.IAMService
}

func (s *IAMServiceServer) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	panic("err")

	token, expireTime, err := s.service.Auth(ctx, req.Email, req.Passwd)
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:    "Access-Token",
		Value:   token,
		Expires: expireTime,
		Path:    "/",
	}
	grpc.SetHeader(ctx, metadata.Pairs("Set-Cookie", cookie.String()))
	grpc.SetHeader(ctx, metadata.Pairs(service.KeyAccessToken, token))
	return &pb.AuthResponse{
		Token:      token,
		ExpireTime: expireTime.Unix(),
	}, nil
}

func (s *IAMServiceServer) CheckAuthStatus(ctx context.Context, request *pb.CheckAuthStatusRequest) (*pb.CheckAuthStatusResponse, error) {
	orgCode, userId, err := s.service.CheckAuthStatus(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	return &pb.CheckAuthStatusResponse{
		OrgCode: orgCode,
		UserId:  userId,
		Status:  true,
	}, nil
}

func (s *IAMServiceServer) CreateAuthority(ctx context.Context, request *pb.CreateAuthorityRequest) (*pb.Authority, error) {
	res, err := s.service.CreateAuthority(ctx, request.Code, request.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Authority{
		Id:    res.ID,
		Code:  res.Code,
		Name:  res.Name,
		Ctime: timestamppb.New(res.Ctime),
		Mtime: timestamppb.New(res.Mtime),
	}, nil
}

func (s *IAMServiceServer) ValidateAuthorities(ctx context.Context, request *pb.ValidateAuthoritiesRequest) (*pb.ValidateAuthoritiesResponse, error) {
	res := &pb.ValidateAuthoritiesResponse{
		Result: make(map[string]bool),
	}

	data, err := s.service.ValidateAuthorities(ctx, request.Token, request.Authorities)
	if err != nil {
		return nil, err
	}

	for _, e := range request.Authorities {
		res.Result[e] = data[e]
	}

	return res, nil
}

func (s *IAMServiceServer) GetAuthority(ctx context.Context, request *pb.GetRequest) (*pb.Authority, error) {
	var isOrgCode bool
	if request.OrgCode != "" {
		isOrgCode = true
	}
	res, err := s.service.GetAuthority(ctx, request.Id, isOrgCode)
	if err != nil {
		return nil, err
	}
	return &pb.Authority{
		Id:    res.ID,
		Code:  res.Code,
		Name:  res.Name,
		Ctime: timestamppb.New(res.Ctime),
		Mtime: timestamppb.New(res.Mtime),
	}, nil
}

func (s *IAMServiceServer) ListAuthority(ctx context.Context, request *pb.ListAuthorityRequest) (*pb.ListAuthorityResponse, error) {
	var (
		isOrgCode          bool
		err                error
		pageToken, groupId *string
	)
	if request.OrgCode != "" {
		isOrgCode = true
	}
	if request.GetPageToken() != nil {
		pageToken = &request.GetPageToken().Value
	}
	if request.GetGroupId() != nil {
		groupId = &request.GetGroupId().Value
	}

	res, err := s.service.ListAuthority(ctx, int(request.GetPageSize()), pageToken, groupId, isOrgCode)
	if err != nil {
		return nil, err
	}
	resItems := res.Items.([]*ent.Authority)
	items := make([]*pb.Authority, 0, len(resItems))
	for _, v := range resItems {
		items = append(items, &pb.Authority{
			Id:    v.ID,
			Code:  v.Code,
			Name:  v.Name,
			Ctime: timestamppb.New(v.Ctime),
			Mtime: timestamppb.New(v.Mtime),
		})
	}
	return &pb.ListAuthorityResponse{
		Pagination: &pb.Pagination{
			PageSize:      uint32(res.PageSize),
			PageTotal:     uint32(res.PageTotal),
			ItemsTotal:    uint32(res.ItemsTotal),
			NextPageToken: res.NextPageToken,
		},
		Items: items,
	}, nil
}

func (s *IAMServiceServer) ListGroup(ctx context.Context, request *pb.ListGroupRequest) (*pb.ListGroupResponse, error) {
	var (
		org              *ent.Org
		err              error
		orgId, pageToken *string
	)
	if request.OrgCode != nil {
		org, err = s.service.GetOrg(ctx, request.GetOrgCode().GetValue())
		if err != nil {
			return nil, err
		}
		orgId = &org.ID
	}
	if request.GetPageToken() != nil {
		pageToken = &request.GetPageToken().Value
	}

	res, err := s.service.ListGroup(ctx, int(request.GetPageSize()), pageToken, orgId)
	if err != nil {
		return nil, err
	}
	resItems := res.Items.([]*ent.Group)
	items := make([]*pb.Group, 0, len(resItems))
	for _, v := range resItems {
		items = append(items, &pb.Group{
			Id:    v.ID,
			Code:  v.Code,
			Name:  v.Name,
			Ctime: timestamppb.New(v.Ctime),
			Mtime: timestamppb.New(v.Mtime),
		})
	}
	return &pb.ListGroupResponse{
		Pagination: &pb.Pagination{
			PageSize:      uint32(res.PageSize),
			PageTotal:     uint32(res.PageTotal),
			ItemsTotal:    uint32(res.ItemsTotal),
			NextPageToken: res.NextPageToken,
		},
		Items: items,
	}, nil
}

func (s *IAMServiceServer) ListUser(ctx context.Context, request *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	var (
		org                             *ent.Org
		err                             error
		orgId, groupId, nsId, pageToken *string
	)
	if request.OrgCode != nil {
		org, err = s.service.GetOrg(ctx, request.GetOrgCode().GetValue())
		if err != nil {
			return nil, err
		}
		orgId = &org.ID
	}
	if request.GetPageToken() != nil {
		pageToken = &request.GetPageToken().Value
	}
	if request.GetGroupId() != nil {
		groupId = &request.GetGroupId().Value
	} else if request.GetNsId() != nil {
		nsId = &request.GetNsId().Value
	}

	res, err := s.service.ListUser(ctx, int(request.GetPageSize()), pageToken, orgId, groupId, nsId)
	if err != nil {
		return nil, err
	}
	resItems := res.Items.([]*ent.User)
	items := make([]*pb.User, 0, len(resItems))
	for _, v := range resItems {
		items = append(items, &pb.User{
			Id:    v.ID,
			Name:  v.Name,
			Email: v.Email,
			Ctime: timestamppb.New(v.Ctime),
			Mtime: timestamppb.New(v.Mtime),
		})
	}
	return &pb.ListUserResponse{
		Pagination: &pb.Pagination{
			PageSize:      uint32(res.PageSize),
			PageTotal:     uint32(res.PageTotal),
			ItemsTotal:    uint32(res.ItemsTotal),
			NextPageToken: res.NextPageToken,
		},
		Items: items,
	}, nil
}

func (s *IAMServiceServer) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest) (*pb.Group, error) {
	res, err := s.service.UpdateGroup(ctx, request.Id, request.Name, request.AddUserIds, request.RemoveUserIds, request.AddAuthorityIds, request.RemoveAuthorityIds)
	if err != nil {
		return nil, err
	}
	return &pb.Group{
		Id:    res.ID,
		Code:  res.Code,
		Name:  res.Name,
		Ctime: timestamppb.New(res.Ctime),
		Mtime: timestamppb.New(res.Mtime),
	}, nil
}

func NewIAMServiceServer(service service.IAMService) *IAMServiceServer {
	return &IAMServiceServer{
		service: service,
	}
}
