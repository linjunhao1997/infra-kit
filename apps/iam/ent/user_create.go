// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"infra-kit/apps/iam/ent/credential"
	"infra-kit/apps/iam/ent/group"
	"infra-kit/apps/iam/ent/namespace"
	"infra-kit/apps/iam/ent/org"
	"infra-kit/apps/iam/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetDisabled sets the "disabled" field.
func (uc *UserCreate) SetDisabled(b bool) *UserCreate {
	uc.mutation.SetDisabled(b)
	return uc
}

// SetVer sets the "ver" field.
func (uc *UserCreate) SetVer(i int) *UserCreate {
	uc.mutation.SetVer(i)
	return uc
}

// SetCtime sets the "ctime" field.
func (uc *UserCreate) SetCtime(t time.Time) *UserCreate {
	uc.mutation.SetCtime(t)
	return uc
}

// SetMtime sets the "mtime" field.
func (uc *UserCreate) SetMtime(t time.Time) *UserCreate {
	uc.mutation.SetMtime(t)
	return uc
}

// SetDeleted sets the "deleted" field.
func (uc *UserCreate) SetDeleted(b bool) *UserCreate {
	uc.mutation.SetDeleted(b)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(s *string) *UserCreate {
	if s != nil {
		uc.SetID(*s)
	}
	return uc
}

// SetOrgID sets the "org" edge to the Org entity by ID.
func (uc *UserCreate) SetOrgID(id string) *UserCreate {
	uc.mutation.SetOrgID(id)
	return uc
}

// SetNillableOrgID sets the "org" edge to the Org entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableOrgID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetOrgID(*id)
	}
	return uc
}

// SetOrg sets the "org" edge to the Org entity.
func (uc *UserCreate) SetOrg(o *Org) *UserCreate {
	return uc.SetOrgID(o.ID)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uc *UserCreate) AddGroupIDs(ids ...string) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the "groups" edges to the Group entity.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddNamespaceIDs adds the "namespace" edge to the Namespace entity by IDs.
func (uc *UserCreate) AddNamespaceIDs(ids ...string) *UserCreate {
	uc.mutation.AddNamespaceIDs(ids...)
	return uc
}

// AddNamespace adds the "namespace" edges to the Namespace entity.
func (uc *UserCreate) AddNamespace(n ...*Namespace) *UserCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uc.AddNamespaceIDs(ids...)
}

// SetCredentialID sets the "credential" edge to the Credential entity by ID.
func (uc *UserCreate) SetCredentialID(id string) *UserCreate {
	uc.mutation.SetCredentialID(id)
	return uc
}

// SetNillableCredentialID sets the "credential" edge to the Credential entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCredentialID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetCredentialID(*id)
	}
	return uc
}

// SetCredential sets the "credential" edge to the Credential entity.
func (uc *UserCreate) SetCredential(c *Credential) *UserCreate {
	return uc.SetCredentialID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if _, ok := uc.mutation.Disabled(); !ok {
		return &ValidationError{Name: "disabled", err: errors.New(`ent: missing required field "User.disabled"`)}
	}
	if _, ok := uc.mutation.Ver(); !ok {
		return &ValidationError{Name: "ver", err: errors.New(`ent: missing required field "User.ver"`)}
	}
	if _, ok := uc.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "User.ctime"`)}
	}
	if _, ok := uc.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New(`ent: missing required field "User.mtime"`)}
	}
	if _, ok := uc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "User.deleted"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Disabled(); ok {
		_spec.SetField(user.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	if value, ok := uc.mutation.Ver(); ok {
		_spec.SetField(user.FieldVer, field.TypeInt, value)
		_node.Ver = value
	}
	if value, ok := uc.mutation.Ctime(); ok {
		_spec.SetField(user.FieldCtime, field.TypeTime, value)
		_node.Ctime = value
	}
	if value, ok := uc.mutation.Mtime(); ok {
		_spec.SetField(user.FieldMtime, field.TypeTime, value)
		_node.Mtime = value
	}
	if value, ok := uc.mutation.Deleted(); ok {
		_spec.SetField(user.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if nodes := uc.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrgTable,
			Columns: []string{user.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.org_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.NamespaceTable,
			Columns: user.NamespacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CredentialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CredentialTable,
			Columns: []string{user.CredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(credential.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsert) SetDisabled(v bool) *UserUpsert {
	u.Set(user.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsert) UpdateDisabled() *UserUpsert {
	u.SetExcluded(user.FieldDisabled)
	return u
}

// SetVer sets the "ver" field.
func (u *UserUpsert) SetVer(v int) *UserUpsert {
	u.Set(user.FieldVer, v)
	return u
}

// UpdateVer sets the "ver" field to the value that was provided on create.
func (u *UserUpsert) UpdateVer() *UserUpsert {
	u.SetExcluded(user.FieldVer)
	return u
}

// AddVer adds v to the "ver" field.
func (u *UserUpsert) AddVer(v int) *UserUpsert {
	u.Add(user.FieldVer, v)
	return u
}

// SetCtime sets the "ctime" field.
func (u *UserUpsert) SetCtime(v time.Time) *UserUpsert {
	u.Set(user.FieldCtime, v)
	return u
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *UserUpsert) UpdateCtime() *UserUpsert {
	u.SetExcluded(user.FieldCtime)
	return u
}

// SetMtime sets the "mtime" field.
func (u *UserUpsert) SetMtime(v time.Time) *UserUpsert {
	u.Set(user.FieldMtime, v)
	return u
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *UserUpsert) UpdateMtime() *UserUpsert {
	u.SetExcluded(user.FieldMtime)
	return u
}

// SetDeleted sets the "deleted" field.
func (u *UserUpsert) SetDeleted(v bool) *UserUpsert {
	u.Set(user.FieldDeleted, v)
	return u
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeleted() *UserUpsert {
	u.SetExcluded(user.FieldDeleted)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsertOne) SetDisabled(v bool) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDisabled() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabled()
	})
}

// SetVer sets the "ver" field.
func (u *UserUpsertOne) SetVer(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetVer(v)
	})
}

// AddVer adds v to the "ver" field.
func (u *UserUpsertOne) AddVer(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddVer(v)
	})
}

// UpdateVer sets the "ver" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateVer() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVer()
	})
}

// SetCtime sets the "ctime" field.
func (u *UserUpsertOne) SetCtime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCtime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCtime()
	})
}

// SetMtime sets the "mtime" field.
func (u *UserUpsertOne) SetMtime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateMtime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateMtime()
	})
}

// SetDeleted sets the "deleted" field.
func (u *UserUpsertOne) SetDeleted(v bool) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeleted(v)
	})
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeleted() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeleted()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetDisabled sets the "disabled" field.
func (u *UserUpsertBulk) SetDisabled(v bool) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDisabled() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDisabled()
	})
}

// SetVer sets the "ver" field.
func (u *UserUpsertBulk) SetVer(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetVer(v)
	})
}

// AddVer adds v to the "ver" field.
func (u *UserUpsertBulk) AddVer(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddVer(v)
	})
}

// UpdateVer sets the "ver" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateVer() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVer()
	})
}

// SetCtime sets the "ctime" field.
func (u *UserUpsertBulk) SetCtime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCtime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCtime()
	})
}

// SetMtime sets the "mtime" field.
func (u *UserUpsertBulk) SetMtime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateMtime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateMtime()
	})
}

// SetDeleted sets the "deleted" field.
func (u *UserUpsertBulk) SetDeleted(v bool) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeleted(v)
	})
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeleted() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeleted()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
