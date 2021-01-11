// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/righttotreatment"
	"github.com/team06/app/ent/righttotreatmenttype"
)

// RightToTreatmentTypeUpdate is the builder for updating RightToTreatmentType entities.
type RightToTreatmentTypeUpdate struct {
	config
	hooks    []Hook
	mutation *RightToTreatmentTypeMutation
}

// Where adds a new predicate for the RightToTreatmentTypeUpdate builder.
func (rtttu *RightToTreatmentTypeUpdate) Where(ps ...predicate.RightToTreatmentType) *RightToTreatmentTypeUpdate {
	rtttu.mutation.predicates = append(rtttu.mutation.predicates, ps...)
	return rtttu
}

// SetTypeName sets the "TypeName" field.
func (rtttu *RightToTreatmentTypeUpdate) SetTypeName(s string) *RightToTreatmentTypeUpdate {
	rtttu.mutation.SetTypeName(s)
	return rtttu
}

// AddTypeIDs adds the "type" edge to the RightToTreatment entity by IDs.
func (rtttu *RightToTreatmentTypeUpdate) AddTypeIDs(ids ...int) *RightToTreatmentTypeUpdate {
	rtttu.mutation.AddTypeIDs(ids...)
	return rtttu
}

// AddType adds the "type" edges to the RightToTreatment entity.
func (rtttu *RightToTreatmentTypeUpdate) AddType(r ...*RightToTreatment) *RightToTreatmentTypeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtttu.AddTypeIDs(ids...)
}

// Mutation returns the RightToTreatmentTypeMutation object of the builder.
func (rtttu *RightToTreatmentTypeUpdate) Mutation() *RightToTreatmentTypeMutation {
	return rtttu.mutation
}

// ClearType clears all "type" edges to the RightToTreatment entity.
func (rtttu *RightToTreatmentTypeUpdate) ClearType() *RightToTreatmentTypeUpdate {
	rtttu.mutation.ClearType()
	return rtttu
}

// RemoveTypeIDs removes the "type" edge to RightToTreatment entities by IDs.
func (rtttu *RightToTreatmentTypeUpdate) RemoveTypeIDs(ids ...int) *RightToTreatmentTypeUpdate {
	rtttu.mutation.RemoveTypeIDs(ids...)
	return rtttu
}

// RemoveType removes "type" edges to RightToTreatment entities.
func (rtttu *RightToTreatmentTypeUpdate) RemoveType(r ...*RightToTreatment) *RightToTreatmentTypeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtttu.RemoveTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtttu *RightToTreatmentTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rtttu.hooks) == 0 {
		if err = rtttu.check(); err != nil {
			return 0, err
		}
		affected, err = rtttu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RightToTreatmentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtttu.check(); err != nil {
				return 0, err
			}
			rtttu.mutation = mutation
			affected, err = rtttu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtttu.hooks) - 1; i >= 0; i-- {
			mut = rtttu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtttu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtttu *RightToTreatmentTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := rtttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtttu *RightToTreatmentTypeUpdate) Exec(ctx context.Context) error {
	_, err := rtttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtttu *RightToTreatmentTypeUpdate) ExecX(ctx context.Context) {
	if err := rtttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtttu *RightToTreatmentTypeUpdate) check() error {
	if v, ok := rtttu.mutation.TypeName(); ok {
		if err := righttotreatmenttype.TypeNameValidator(v); err != nil {
			return &ValidationError{Name: "TypeName", err: fmt.Errorf("ent: validator failed for field \"TypeName\": %w", err)}
		}
	}
	return nil
}

func (rtttu *RightToTreatmentTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   righttotreatmenttype.Table,
			Columns: righttotreatmenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: righttotreatmenttype.FieldID,
			},
		},
	}
	if ps := rtttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtttu.mutation.TypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: righttotreatmenttype.FieldTypeName,
		})
	}
	if rtttu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtttu.mutation.RemovedTypeIDs(); len(nodes) > 0 && !rtttu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtttu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{righttotreatmenttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RightToTreatmentTypeUpdateOne is the builder for updating a single RightToTreatmentType entity.
type RightToTreatmentTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *RightToTreatmentTypeMutation
}

// SetTypeName sets the "TypeName" field.
func (rtttuo *RightToTreatmentTypeUpdateOne) SetTypeName(s string) *RightToTreatmentTypeUpdateOne {
	rtttuo.mutation.SetTypeName(s)
	return rtttuo
}

// AddTypeIDs adds the "type" edge to the RightToTreatment entity by IDs.
func (rtttuo *RightToTreatmentTypeUpdateOne) AddTypeIDs(ids ...int) *RightToTreatmentTypeUpdateOne {
	rtttuo.mutation.AddTypeIDs(ids...)
	return rtttuo
}

// AddType adds the "type" edges to the RightToTreatment entity.
func (rtttuo *RightToTreatmentTypeUpdateOne) AddType(r ...*RightToTreatment) *RightToTreatmentTypeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtttuo.AddTypeIDs(ids...)
}

// Mutation returns the RightToTreatmentTypeMutation object of the builder.
func (rtttuo *RightToTreatmentTypeUpdateOne) Mutation() *RightToTreatmentTypeMutation {
	return rtttuo.mutation
}

// ClearType clears all "type" edges to the RightToTreatment entity.
func (rtttuo *RightToTreatmentTypeUpdateOne) ClearType() *RightToTreatmentTypeUpdateOne {
	rtttuo.mutation.ClearType()
	return rtttuo
}

// RemoveTypeIDs removes the "type" edge to RightToTreatment entities by IDs.
func (rtttuo *RightToTreatmentTypeUpdateOne) RemoveTypeIDs(ids ...int) *RightToTreatmentTypeUpdateOne {
	rtttuo.mutation.RemoveTypeIDs(ids...)
	return rtttuo
}

// RemoveType removes "type" edges to RightToTreatment entities.
func (rtttuo *RightToTreatmentTypeUpdateOne) RemoveType(r ...*RightToTreatment) *RightToTreatmentTypeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtttuo.RemoveTypeIDs(ids...)
}

// Save executes the query and returns the updated RightToTreatmentType entity.
func (rtttuo *RightToTreatmentTypeUpdateOne) Save(ctx context.Context) (*RightToTreatmentType, error) {
	var (
		err  error
		node *RightToTreatmentType
	)
	if len(rtttuo.hooks) == 0 {
		if err = rtttuo.check(); err != nil {
			return nil, err
		}
		node, err = rtttuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RightToTreatmentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtttuo.check(); err != nil {
				return nil, err
			}
			rtttuo.mutation = mutation
			node, err = rtttuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtttuo.hooks) - 1; i >= 0; i-- {
			mut = rtttuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtttuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtttuo *RightToTreatmentTypeUpdateOne) SaveX(ctx context.Context) *RightToTreatmentType {
	node, err := rtttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtttuo *RightToTreatmentTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := rtttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtttuo *RightToTreatmentTypeUpdateOne) ExecX(ctx context.Context) {
	if err := rtttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtttuo *RightToTreatmentTypeUpdateOne) check() error {
	if v, ok := rtttuo.mutation.TypeName(); ok {
		if err := righttotreatmenttype.TypeNameValidator(v); err != nil {
			return &ValidationError{Name: "TypeName", err: fmt.Errorf("ent: validator failed for field \"TypeName\": %w", err)}
		}
	}
	return nil
}

func (rtttuo *RightToTreatmentTypeUpdateOne) sqlSave(ctx context.Context) (_node *RightToTreatmentType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   righttotreatmenttype.Table,
			Columns: righttotreatmenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: righttotreatmenttype.FieldID,
			},
		},
	}
	id, ok := rtttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RightToTreatmentType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := rtttuo.mutation.TypeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: righttotreatmenttype.FieldTypeName,
		})
	}
	if rtttuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtttuo.mutation.RemovedTypeIDs(); len(nodes) > 0 && !rtttuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtttuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   righttotreatmenttype.TypeTable,
			Columns: []string{righttotreatmenttype.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: righttotreatment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RightToTreatmentType{config: rtttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{righttotreatmenttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
