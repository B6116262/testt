// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/proceduretype"
)

// ProcedureTypeDelete is the builder for deleting a ProcedureType entity.
type ProcedureTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProcedureTypeMutation
}

// Where adds a new predicate to the ProcedureTypeDelete builder.
func (ptd *ProcedureTypeDelete) Where(ps ...predicate.ProcedureType) *ProcedureTypeDelete {
	ptd.mutation.predicates = append(ptd.mutation.predicates, ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *ProcedureTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptd.hooks) == 0 {
		affected, err = ptd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProcedureTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptd.mutation = mutation
			affected, err = ptd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptd.hooks) - 1; i >= 0; i-- {
			mut = ptd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *ProcedureTypeDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *ProcedureTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: proceduretype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: proceduretype.FieldID,
			},
		},
	}
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
}

// ProcedureTypeDeleteOne is the builder for deleting a single ProcedureType entity.
type ProcedureTypeDeleteOne struct {
	ptd *ProcedureTypeDelete
}

// Exec executes the deletion query.
func (ptdo *ProcedureTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{proceduretype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *ProcedureTypeDeleteOne) ExecX(ctx context.Context) {
	ptdo.ptd.ExecX(ctx)
}