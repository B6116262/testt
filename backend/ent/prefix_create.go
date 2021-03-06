// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/prefix"
)

// PrefixCreate is the builder for creating a Prefix entity.
type PrefixCreate struct {
	config
	mutation *PrefixMutation
	hooks    []Hook
}

// SetPrefix sets the "prefix" field.
func (pc *PrefixCreate) SetPrefix(s string) *PrefixCreate {
	pc.mutation.SetPrefix(s)
	return pc
}

// AddPrefixToPatientIDs adds the "PrefixToPatient" edge to the Patient entity by IDs.
func (pc *PrefixCreate) AddPrefixToPatientIDs(ids ...int) *PrefixCreate {
	pc.mutation.AddPrefixToPatientIDs(ids...)
	return pc
}

// AddPrefixToPatient adds the "PrefixToPatient" edges to the Patient entity.
func (pc *PrefixCreate) AddPrefixToPatient(p ...*Patient) *PrefixCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPrefixToPatientIDs(ids...)
}

// Mutation returns the PrefixMutation object of the builder.
func (pc *PrefixCreate) Mutation() *PrefixMutation {
	return pc.mutation
}

// Save creates the Prefix in the database.
func (pc *PrefixCreate) Save(ctx context.Context) (*Prefix, error) {
	var (
		err  error
		node *Prefix
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrefixMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrefixCreate) SaveX(ctx context.Context) *Prefix {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PrefixCreate) check() error {
	if _, ok := pc.mutation.Prefix(); !ok {
		return &ValidationError{Name: "prefix", err: errors.New("ent: missing required field \"prefix\"")}
	}
	if v, ok := pc.mutation.Prefix(); ok {
		if err := prefix.PrefixValidator(v); err != nil {
			return &ValidationError{Name: "prefix", err: fmt.Errorf("ent: validator failed for field \"prefix\": %w", err)}
		}
	}
	return nil
}

func (pc *PrefixCreate) sqlSave(ctx context.Context) (*Prefix, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PrefixCreate) createSpec() (*Prefix, *sqlgraph.CreateSpec) {
	var (
		_node = &Prefix{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prefix.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prefix.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Prefix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prefix.FieldPrefix,
		})
		_node.Prefix = value
	}
	if nodes := pc.mutation.PrefixToPatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.PrefixToPatientTable,
			Columns: []string{prefix.PrefixToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
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

// PrefixCreateBulk is the builder for creating many Prefix entities in bulk.
type PrefixCreateBulk struct {
	config
	builders []*PrefixCreate
}

// Save creates the Prefix entities in the database.
func (pcb *PrefixCreateBulk) Save(ctx context.Context) ([]*Prefix, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Prefix, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrefixMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PrefixCreateBulk) SaveX(ctx context.Context) []*Prefix {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
