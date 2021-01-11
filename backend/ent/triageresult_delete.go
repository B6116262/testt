// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/triageresult"
)

// TriageResultDelete is the builder for deleting a TriageResult entity.
type TriageResultDelete struct {
	config
	hooks    []Hook
	mutation *TriageResultMutation
}

// Where adds a new predicate to the TriageResultDelete builder.
func (trd *TriageResultDelete) Where(ps ...predicate.TriageResult) *TriageResultDelete {
	trd.mutation.predicates = append(trd.mutation.predicates, ps...)
	return trd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (trd *TriageResultDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(trd.hooks) == 0 {
		affected, err = trd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TriageResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			trd.mutation = mutation
			affected, err = trd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(trd.hooks) - 1; i >= 0; i-- {
			mut = trd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (trd *TriageResultDelete) ExecX(ctx context.Context) int {
	n, err := trd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (trd *TriageResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: triageresult.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: triageresult.FieldID,
			},
		},
	}
	if ps := trd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, trd.driver, _spec)
}

// TriageResultDeleteOne is the builder for deleting a single TriageResult entity.
type TriageResultDeleteOne struct {
	trd *TriageResultDelete
}

// Exec executes the deletion query.
func (trdo *TriageResultDeleteOne) Exec(ctx context.Context) error {
	n, err := trdo.trd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{triageresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (trdo *TriageResultDeleteOne) ExecX(ctx context.Context) {
	trdo.trd.ExecX(ctx)
}