// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"infra-kit/apps/iam/ent/authority"
	"infra-kit/apps/iam/ent/group"
	"infra-kit/apps/iam/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthorityUpdate is the builder for updating Authority entities.
type AuthorityUpdate struct {
	config
	hooks    []Hook
	mutation *AuthorityMutation
}

// Where appends a list predicates to the AuthorityUpdate builder.
func (au *AuthorityUpdate) Where(ps ...predicate.Authority) *AuthorityUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCode sets the "code" field.
func (au *AuthorityUpdate) SetCode(s string) *AuthorityUpdate {
	au.mutation.SetCode(s)
	return au
}

// SetName sets the "name" field.
func (au *AuthorityUpdate) SetName(s string) *AuthorityUpdate {
	au.mutation.SetName(s)
	return au
}

// SetCtime sets the "ctime" field.
func (au *AuthorityUpdate) SetCtime(t time.Time) *AuthorityUpdate {
	au.mutation.SetCtime(t)
	return au
}

// SetMtime sets the "mtime" field.
func (au *AuthorityUpdate) SetMtime(t time.Time) *AuthorityUpdate {
	au.mutation.SetMtime(t)
	return au
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (au *AuthorityUpdate) AddGroupIDs(ids ...string) *AuthorityUpdate {
	au.mutation.AddGroupIDs(ids...)
	return au
}

// AddGroups adds the "groups" edges to the Group entity.
func (au *AuthorityUpdate) AddGroups(g ...*Group) *AuthorityUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.AddGroupIDs(ids...)
}

// Mutation returns the AuthorityMutation object of the builder.
func (au *AuthorityUpdate) Mutation() *AuthorityMutation {
	return au.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (au *AuthorityUpdate) ClearGroups() *AuthorityUpdate {
	au.mutation.ClearGroups()
	return au
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (au *AuthorityUpdate) RemoveGroupIDs(ids ...string) *AuthorityUpdate {
	au.mutation.RemoveGroupIDs(ids...)
	return au
}

// RemoveGroups removes "groups" edges to Group entities.
func (au *AuthorityUpdate) RemoveGroups(g ...*Group) *AuthorityUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthorityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthorityUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthorityUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthorityUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AuthorityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(authority.Table, authority.Columns, sqlgraph.NewFieldSpec(authority.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Code(); ok {
		_spec.SetField(authority.FieldCode, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(authority.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Ctime(); ok {
		_spec.SetField(authority.FieldCtime, field.TypeTime, value)
	}
	if value, ok := au.mutation.Mtime(); ok {
		_spec.SetField(authority.FieldMtime, field.TypeTime, value)
	}
	if au.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !au.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authority.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AuthorityUpdateOne is the builder for updating a single Authority entity.
type AuthorityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthorityMutation
}

// SetCode sets the "code" field.
func (auo *AuthorityUpdateOne) SetCode(s string) *AuthorityUpdateOne {
	auo.mutation.SetCode(s)
	return auo
}

// SetName sets the "name" field.
func (auo *AuthorityUpdateOne) SetName(s string) *AuthorityUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetCtime sets the "ctime" field.
func (auo *AuthorityUpdateOne) SetCtime(t time.Time) *AuthorityUpdateOne {
	auo.mutation.SetCtime(t)
	return auo
}

// SetMtime sets the "mtime" field.
func (auo *AuthorityUpdateOne) SetMtime(t time.Time) *AuthorityUpdateOne {
	auo.mutation.SetMtime(t)
	return auo
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (auo *AuthorityUpdateOne) AddGroupIDs(ids ...string) *AuthorityUpdateOne {
	auo.mutation.AddGroupIDs(ids...)
	return auo
}

// AddGroups adds the "groups" edges to the Group entity.
func (auo *AuthorityUpdateOne) AddGroups(g ...*Group) *AuthorityUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.AddGroupIDs(ids...)
}

// Mutation returns the AuthorityMutation object of the builder.
func (auo *AuthorityUpdateOne) Mutation() *AuthorityMutation {
	return auo.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (auo *AuthorityUpdateOne) ClearGroups() *AuthorityUpdateOne {
	auo.mutation.ClearGroups()
	return auo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (auo *AuthorityUpdateOne) RemoveGroupIDs(ids ...string) *AuthorityUpdateOne {
	auo.mutation.RemoveGroupIDs(ids...)
	return auo
}

// RemoveGroups removes "groups" edges to Group entities.
func (auo *AuthorityUpdateOne) RemoveGroups(g ...*Group) *AuthorityUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the AuthorityUpdate builder.
func (auo *AuthorityUpdateOne) Where(ps ...predicate.Authority) *AuthorityUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthorityUpdateOne) Select(field string, fields ...string) *AuthorityUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Authority entity.
func (auo *AuthorityUpdateOne) Save(ctx context.Context) (*Authority, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthorityUpdateOne) SaveX(ctx context.Context) *Authority {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthorityUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthorityUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AuthorityUpdateOne) sqlSave(ctx context.Context) (_node *Authority, err error) {
	_spec := sqlgraph.NewUpdateSpec(authority.Table, authority.Columns, sqlgraph.NewFieldSpec(authority.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Authority.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authority.FieldID)
		for _, f := range fields {
			if !authority.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authority.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Code(); ok {
		_spec.SetField(authority.FieldCode, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(authority.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Ctime(); ok {
		_spec.SetField(authority.FieldCtime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Mtime(); ok {
		_spec.SetField(authority.FieldMtime, field.TypeTime, value)
	}
	if auo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !auo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   authority.GroupsTable,
			Columns: authority.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Authority{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authority.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
