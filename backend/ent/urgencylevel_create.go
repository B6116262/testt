// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/triageresult"
	"github.com/team06/app/ent/urgencylevel"
)

// UrgencyLevelCreate is the builder for creating a UrgencyLevel entity.
type UrgencyLevelCreate struct {
	config
	mutation *UrgencyLevelMutation
	hooks    []Hook
}

// SetUrgencyName sets the "urgencyName" field.
func (ulc *UrgencyLevelCreate) SetUrgencyName(s string) *UrgencyLevelCreate {
	ulc.mutation.SetUrgencyName(s)
	return ulc
}

// AddUrgencyLevelToTriageResultIDs adds the "urgencyLevelToTriageResult" edge to the TriageResult entity by IDs.
func (ulc *UrgencyLevelCreate) AddUrgencyLevelToTriageResultIDs(ids ...int) *UrgencyLevelCreate {
	ulc.mutation.AddUrgencyLevelToTriageResultIDs(ids...)
	return ulc
}

// AddUrgencyLevelToTriageResult adds the "urgencyLevelToTriageResult" edges to the TriageResult entity.
func (ulc *UrgencyLevelCreate) AddUrgencyLevelToTriageResult(t ...*TriageResult) *UrgencyLevelCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ulc.AddUrgencyLevelToTriageResultIDs(ids...)
}

// Mutation returns the UrgencyLevelMutation object of the builder.
func (ulc *UrgencyLevelCreate) Mutation() *UrgencyLevelMutation {
	return ulc.mutation
}

// Save creates the UrgencyLevel in the database.
func (ulc *UrgencyLevelCreate) Save(ctx context.Context) (*UrgencyLevel, error) {
	var (
		err  error
		node *UrgencyLevel
	)
	if len(ulc.hooks) == 0 {
		if err = ulc.check(); err != nil {
			return nil, err
		}
		node, err = ulc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UrgencyLevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ulc.check(); err != nil {
				return nil, err
			}
			ulc.mutation = mutation
			node, err = ulc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ulc.hooks) - 1; i >= 0; i-- {
			mut = ulc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ulc *UrgencyLevelCreate) SaveX(ctx context.Context) *UrgencyLevel {
	v, err := ulc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ulc *UrgencyLevelCreate) check() error {
	if _, ok := ulc.mutation.UrgencyName(); !ok {
		return &ValidationError{Name: "urgencyName", err: errors.New("ent: missing required field \"urgencyName\"")}
	}
	if v, ok := ulc.mutation.UrgencyName(); ok {
		if err := urgencylevel.UrgencyNameValidator(v); err != nil {
			return &ValidationError{Name: "urgencyName", err: fmt.Errorf("ent: validator failed for field \"urgencyName\": %w", err)}
		}
	}
	return nil
}

func (ulc *UrgencyLevelCreate) sqlSave(ctx context.Context) (*UrgencyLevel, error) {
	_node, _spec := ulc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ulc *UrgencyLevelCreate) createSpec() (*UrgencyLevel, *sqlgraph.CreateSpec) {
	var (
		_node = &UrgencyLevel{config: ulc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: urgencylevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: urgencylevel.FieldID,
			},
		}
	)
	if value, ok := ulc.mutation.UrgencyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urgencylevel.FieldUrgencyName,
		})
		_node.UrgencyName = value
	}
	if nodes := ulc.mutation.UrgencyLevelToTriageResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   urgencylevel.UrgencyLevelToTriageResultTable,
			Columns: []string{urgencylevel.UrgencyLevelToTriageResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: triageresult.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UrgencyLevelCreateBulk is the builder for creating many UrgencyLevel entities in bulk.
type UrgencyLevelCreateBulk struct {
	config
	builders []*UrgencyLevelCreate
}

// Save creates the UrgencyLevel entities in the database.
func (ulcb *UrgencyLevelCreateBulk) Save(ctx context.Context) ([]*UrgencyLevel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ulcb.builders))
	nodes := make([]*UrgencyLevel, len(ulcb.builders))
	mutators := make([]Mutator, len(ulcb.builders))
	for i := range ulcb.builders {
		func(i int, root context.Context) {
			builder := ulcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UrgencyLevelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ulcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ulcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulcb *UrgencyLevelCreateBulk) SaveX(ctx context.Context) []*UrgencyLevel {
	v, err := ulcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
