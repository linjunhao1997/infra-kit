// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"infra-kit/apps/iam/ent/namespace"
	"infra-kit/apps/iam/ent/org"
	"infra-kit/apps/iam/ent/predicate"
	"infra-kit/apps/iam/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NamespaceUpdate is the builder for updating Namespace entities.
type NamespaceUpdate struct {
	config
	hooks    []Hook
	mutation *NamespaceMutation
}

// Where appends a list predicates to the NamespaceUpdate builder.
func (nu *NamespaceUpdate) Where(ps ...predicate.Namespace) *NamespaceUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetCode sets the "code" field.
func (nu *NamespaceUpdate) SetCode(s string) *NamespaceUpdate {
	nu.mutation.SetCode(s)
	return nu
}

// SetName sets the "name" field.
func (nu *NamespaceUpdate) SetName(s string) *NamespaceUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetDescription sets the "description" field.
func (nu *NamespaceUpdate) SetDescription(s string) *NamespaceUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetLabel sets the "label" field.
func (nu *NamespaceUpdate) SetLabel(s string) *NamespaceUpdate {
	nu.mutation.SetLabel(s)
	return nu
}

// SetCtime sets the "ctime" field.
func (nu *NamespaceUpdate) SetCtime(t time.Time) *NamespaceUpdate {
	nu.mutation.SetCtime(t)
	return nu
}

// SetMtime sets the "mtime" field.
func (nu *NamespaceUpdate) SetMtime(t time.Time) *NamespaceUpdate {
	nu.mutation.SetMtime(t)
	return nu
}

// SetDeleted sets the "deleted" field.
func (nu *NamespaceUpdate) SetDeleted(b bool) *NamespaceUpdate {
	nu.mutation.SetDeleted(b)
	return nu
}

// SetOrgID sets the "org" edge to the Org entity by ID.
func (nu *NamespaceUpdate) SetOrgID(id string) *NamespaceUpdate {
	nu.mutation.SetOrgID(id)
	return nu
}

// SetNillableOrgID sets the "org" edge to the Org entity by ID if the given value is not nil.
func (nu *NamespaceUpdate) SetNillableOrgID(id *string) *NamespaceUpdate {
	if id != nil {
		nu = nu.SetOrgID(*id)
	}
	return nu
}

// SetOrg sets the "org" edge to the Org entity.
func (nu *NamespaceUpdate) SetOrg(o *Org) *NamespaceUpdate {
	return nu.SetOrgID(o.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (nu *NamespaceUpdate) AddUserIDs(ids ...string) *NamespaceUpdate {
	nu.mutation.AddUserIDs(ids...)
	return nu
}

// AddUsers adds the "users" edges to the User entity.
func (nu *NamespaceUpdate) AddUsers(u ...*User) *NamespaceUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.AddUserIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nu *NamespaceUpdate) Mutation() *NamespaceMutation {
	return nu.mutation
}

// ClearOrg clears the "org" edge to the Org entity.
func (nu *NamespaceUpdate) ClearOrg() *NamespaceUpdate {
	nu.mutation.ClearOrg()
	return nu
}

// ClearUsers clears all "users" edges to the User entity.
func (nu *NamespaceUpdate) ClearUsers() *NamespaceUpdate {
	nu.mutation.ClearUsers()
	return nu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (nu *NamespaceUpdate) RemoveUserIDs(ids ...string) *NamespaceUpdate {
	nu.mutation.RemoveUserIDs(ids...)
	return nu
}

// RemoveUsers removes "users" edges to User entities.
func (nu *NamespaceUpdate) RemoveUsers(u ...*User) *NamespaceUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NamespaceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NamespaceUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NamespaceUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NamespaceUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NamespaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(namespace.Table, namespace.Columns, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeString))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Code(); ok {
		_spec.SetField(namespace.FieldCode, field.TypeString, value)
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(namespace.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.SetField(namespace.FieldDescription, field.TypeString, value)
	}
	if value, ok := nu.mutation.Label(); ok {
		_spec.SetField(namespace.FieldLabel, field.TypeString, value)
	}
	if value, ok := nu.mutation.Ctime(); ok {
		_spec.SetField(namespace.FieldCtime, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Mtime(); ok {
		_spec.SetField(namespace.FieldMtime, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Deleted(); ok {
		_spec.SetField(namespace.FieldDeleted, field.TypeBool, value)
	}
	if nu.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   namespace.OrgTable,
			Columns: []string{namespace.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   namespace.OrgTable,
			Columns: []string{namespace.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !nu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NamespaceUpdateOne is the builder for updating a single Namespace entity.
type NamespaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NamespaceMutation
}

// SetCode sets the "code" field.
func (nuo *NamespaceUpdateOne) SetCode(s string) *NamespaceUpdateOne {
	nuo.mutation.SetCode(s)
	return nuo
}

// SetName sets the "name" field.
func (nuo *NamespaceUpdateOne) SetName(s string) *NamespaceUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NamespaceUpdateOne) SetDescription(s string) *NamespaceUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetLabel sets the "label" field.
func (nuo *NamespaceUpdateOne) SetLabel(s string) *NamespaceUpdateOne {
	nuo.mutation.SetLabel(s)
	return nuo
}

// SetCtime sets the "ctime" field.
func (nuo *NamespaceUpdateOne) SetCtime(t time.Time) *NamespaceUpdateOne {
	nuo.mutation.SetCtime(t)
	return nuo
}

// SetMtime sets the "mtime" field.
func (nuo *NamespaceUpdateOne) SetMtime(t time.Time) *NamespaceUpdateOne {
	nuo.mutation.SetMtime(t)
	return nuo
}

// SetDeleted sets the "deleted" field.
func (nuo *NamespaceUpdateOne) SetDeleted(b bool) *NamespaceUpdateOne {
	nuo.mutation.SetDeleted(b)
	return nuo
}

// SetOrgID sets the "org" edge to the Org entity by ID.
func (nuo *NamespaceUpdateOne) SetOrgID(id string) *NamespaceUpdateOne {
	nuo.mutation.SetOrgID(id)
	return nuo
}

// SetNillableOrgID sets the "org" edge to the Org entity by ID if the given value is not nil.
func (nuo *NamespaceUpdateOne) SetNillableOrgID(id *string) *NamespaceUpdateOne {
	if id != nil {
		nuo = nuo.SetOrgID(*id)
	}
	return nuo
}

// SetOrg sets the "org" edge to the Org entity.
func (nuo *NamespaceUpdateOne) SetOrg(o *Org) *NamespaceUpdateOne {
	return nuo.SetOrgID(o.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (nuo *NamespaceUpdateOne) AddUserIDs(ids ...string) *NamespaceUpdateOne {
	nuo.mutation.AddUserIDs(ids...)
	return nuo
}

// AddUsers adds the "users" edges to the User entity.
func (nuo *NamespaceUpdateOne) AddUsers(u ...*User) *NamespaceUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.AddUserIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nuo *NamespaceUpdateOne) Mutation() *NamespaceMutation {
	return nuo.mutation
}

// ClearOrg clears the "org" edge to the Org entity.
func (nuo *NamespaceUpdateOne) ClearOrg() *NamespaceUpdateOne {
	nuo.mutation.ClearOrg()
	return nuo
}

// ClearUsers clears all "users" edges to the User entity.
func (nuo *NamespaceUpdateOne) ClearUsers() *NamespaceUpdateOne {
	nuo.mutation.ClearUsers()
	return nuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (nuo *NamespaceUpdateOne) RemoveUserIDs(ids ...string) *NamespaceUpdateOne {
	nuo.mutation.RemoveUserIDs(ids...)
	return nuo
}

// RemoveUsers removes "users" edges to User entities.
func (nuo *NamespaceUpdateOne) RemoveUsers(u ...*User) *NamespaceUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the NamespaceUpdate builder.
func (nuo *NamespaceUpdateOne) Where(ps ...predicate.Namespace) *NamespaceUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NamespaceUpdateOne) Select(field string, fields ...string) *NamespaceUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Namespace entity.
func (nuo *NamespaceUpdateOne) Save(ctx context.Context) (*Namespace, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NamespaceUpdateOne) SaveX(ctx context.Context) *Namespace {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NamespaceUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NamespaceUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NamespaceUpdateOne) sqlSave(ctx context.Context) (_node *Namespace, err error) {
	_spec := sqlgraph.NewUpdateSpec(namespace.Table, namespace.Columns, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeString))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Namespace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namespace.FieldID)
		for _, f := range fields {
			if !namespace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != namespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Code(); ok {
		_spec.SetField(namespace.FieldCode, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(namespace.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.SetField(namespace.FieldDescription, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Label(); ok {
		_spec.SetField(namespace.FieldLabel, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Ctime(); ok {
		_spec.SetField(namespace.FieldCtime, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Mtime(); ok {
		_spec.SetField(namespace.FieldMtime, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Deleted(); ok {
		_spec.SetField(namespace.FieldDeleted, field.TypeBool, value)
	}
	if nuo.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   namespace.OrgTable,
			Columns: []string{namespace.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   namespace.OrgTable,
			Columns: []string{namespace.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !nuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   namespace.UsersTable,
			Columns: namespace.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Namespace{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
