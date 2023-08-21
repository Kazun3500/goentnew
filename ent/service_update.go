// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"goentnew/ent/predicate"
	"goentnew/ent/service"
	"goentnew/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceUpdate is the builder for updating Service entities.
type ServiceUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceMutation
}

// Where appends a list predicates to the ServiceUpdate builder.
func (su *ServiceUpdate) Where(ps ...predicate.Service) *ServiceUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *ServiceUpdate) SetName(s string) *ServiceUpdate {
	su.mutation.SetName(s)
	return su
}

// SetType sets the "type" field.
func (su *ServiceUpdate) SetType(s service.Type) *ServiceUpdate {
	su.mutation.SetType(s)
	return su
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (su *ServiceUpdate) SetOwnerID(id int) *ServiceUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (su *ServiceUpdate) SetNillableOwnerID(id *int) *ServiceUpdate {
	if id != nil {
		su = su.SetOwnerID(*id)
	}
	return su
}

// SetOwner sets the "owner" edge to the User entity.
func (su *ServiceUpdate) SetOwner(u *User) *ServiceUpdate {
	return su.SetOwnerID(u.ID)
}

// Mutation returns the ServiceMutation object of the builder.
func (su *ServiceUpdate) Mutation() *ServiceMutation {
	return su.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (su *ServiceUpdate) ClearOwner() *ServiceUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ServiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServiceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServiceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ServiceUpdate) check() error {
	if v, ok := su.mutation.GetType(); ok {
		if err := service.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Service.type": %w`, err)}
		}
	}
	return nil
}

func (su *ServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(service.FieldType, field.TypeEnum, value)
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.OwnerTable,
			Columns: []string{service.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.OwnerTable,
			Columns: []string{service.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ServiceUpdateOne is the builder for updating a single Service entity.
type ServiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceMutation
}

// SetName sets the "name" field.
func (suo *ServiceUpdateOne) SetName(s string) *ServiceUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetType sets the "type" field.
func (suo *ServiceUpdateOne) SetType(s service.Type) *ServiceUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (suo *ServiceUpdateOne) SetOwnerID(id int) *ServiceUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableOwnerID(id *int) *ServiceUpdateOne {
	if id != nil {
		suo = suo.SetOwnerID(*id)
	}
	return suo
}

// SetOwner sets the "owner" edge to the User entity.
func (suo *ServiceUpdateOne) SetOwner(u *User) *ServiceUpdateOne {
	return suo.SetOwnerID(u.ID)
}

// Mutation returns the ServiceMutation object of the builder.
func (suo *ServiceUpdateOne) Mutation() *ServiceMutation {
	return suo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (suo *ServiceUpdateOne) ClearOwner() *ServiceUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Where appends a list predicates to the ServiceUpdate builder.
func (suo *ServiceUpdateOne) Where(ps ...predicate.Service) *ServiceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ServiceUpdateOne) Select(field string, fields ...string) *ServiceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Service entity.
func (suo *ServiceUpdateOne) Save(ctx context.Context) (*Service, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServiceUpdateOne) SaveX(ctx context.Context) *Service {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServiceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ServiceUpdateOne) check() error {
	if v, ok := suo.mutation.GetType(); ok {
		if err := service.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Service.type": %w`, err)}
		}
	}
	return nil
}

func (suo *ServiceUpdateOne) sqlSave(ctx context.Context) (_node *Service, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(service.Table, service.Columns, sqlgraph.NewFieldSpec(service.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Service.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, service.FieldID)
		for _, f := range fields {
			if !service.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != service.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(service.FieldType, field.TypeEnum, value)
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.OwnerTable,
			Columns: []string{service.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.OwnerTable,
			Columns: []string{service.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Service{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{service.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}