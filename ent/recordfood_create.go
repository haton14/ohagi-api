// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/haton14/ohagi-api/ent/recordfood"
)

// RecordFoodCreate is the builder for creating a RecordFood entity.
type RecordFoodCreate struct {
	config
	mutation *RecordFoodMutation
	hooks    []Hook
}

// SetRecordID sets the "record_id" field.
func (rfc *RecordFoodCreate) SetRecordID(i int) *RecordFoodCreate {
	rfc.mutation.SetRecordID(i)
	return rfc
}

// SetFoodID sets the "food_id" field.
func (rfc *RecordFoodCreate) SetFoodID(i int) *RecordFoodCreate {
	rfc.mutation.SetFoodID(i)
	return rfc
}

// SetAmount sets the "amount" field.
func (rfc *RecordFoodCreate) SetAmount(f float64) *RecordFoodCreate {
	rfc.mutation.SetAmount(f)
	return rfc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (rfc *RecordFoodCreate) SetNillableAmount(f *float64) *RecordFoodCreate {
	if f != nil {
		rfc.SetAmount(*f)
	}
	return rfc
}

// Mutation returns the RecordFoodMutation object of the builder.
func (rfc *RecordFoodCreate) Mutation() *RecordFoodMutation {
	return rfc.mutation
}

// Save creates the RecordFood in the database.
func (rfc *RecordFoodCreate) Save(ctx context.Context) (*RecordFood, error) {
	var (
		err  error
		node *RecordFood
	)
	rfc.defaults()
	if len(rfc.hooks) == 0 {
		if err = rfc.check(); err != nil {
			return nil, err
		}
		node, err = rfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordFoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rfc.check(); err != nil {
				return nil, err
			}
			rfc.mutation = mutation
			if node, err = rfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rfc.hooks) - 1; i >= 0; i-- {
			if rfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rfc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rfc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rfc *RecordFoodCreate) SaveX(ctx context.Context) *RecordFood {
	v, err := rfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rfc *RecordFoodCreate) Exec(ctx context.Context) error {
	_, err := rfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfc *RecordFoodCreate) ExecX(ctx context.Context) {
	if err := rfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rfc *RecordFoodCreate) defaults() {
	if _, ok := rfc.mutation.Amount(); !ok {
		v := recordfood.DefaultAmount
		rfc.mutation.SetAmount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rfc *RecordFoodCreate) check() error {
	if _, ok := rfc.mutation.RecordID(); !ok {
		return &ValidationError{Name: "record_id", err: errors.New(`ent: missing required field "record_id"`)}
	}
	if v, ok := rfc.mutation.RecordID(); ok {
		if err := recordfood.RecordIDValidator(v); err != nil {
			return &ValidationError{Name: "record_id", err: fmt.Errorf(`ent: validator failed for field "record_id": %w`, err)}
		}
	}
	if _, ok := rfc.mutation.FoodID(); !ok {
		return &ValidationError{Name: "food_id", err: errors.New(`ent: missing required field "food_id"`)}
	}
	if v, ok := rfc.mutation.FoodID(); ok {
		if err := recordfood.FoodIDValidator(v); err != nil {
			return &ValidationError{Name: "food_id", err: fmt.Errorf(`ent: validator failed for field "food_id": %w`, err)}
		}
	}
	if _, ok := rfc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if v, ok := rfc.mutation.Amount(); ok {
		if err := recordfood.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "amount": %w`, err)}
		}
	}
	return nil
}

func (rfc *RecordFoodCreate) sqlSave(ctx context.Context) (*RecordFood, error) {
	_node, _spec := rfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rfc *RecordFoodCreate) createSpec() (*RecordFood, *sqlgraph.CreateSpec) {
	var (
		_node = &RecordFood{config: rfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recordfood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recordfood.FieldID,
			},
		}
	)
	if value, ok := rfc.mutation.RecordID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recordfood.FieldRecordID,
		})
		_node.RecordID = value
	}
	if value, ok := rfc.mutation.FoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recordfood.FieldFoodID,
		})
		_node.FoodID = value
	}
	if value, ok := rfc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: recordfood.FieldAmount,
		})
		_node.Amount = value
	}
	return _node, _spec
}

// RecordFoodCreateBulk is the builder for creating many RecordFood entities in bulk.
type RecordFoodCreateBulk struct {
	config
	builders []*RecordFoodCreate
}

// Save creates the RecordFood entities in the database.
func (rfcb *RecordFoodCreateBulk) Save(ctx context.Context) ([]*RecordFood, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rfcb.builders))
	nodes := make([]*RecordFood, len(rfcb.builders))
	mutators := make([]Mutator, len(rfcb.builders))
	for i := range rfcb.builders {
		func(i int, root context.Context) {
			builder := rfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordFoodMutation)
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
					_, err = mutators[i+1].Mutate(root, rfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rfcb *RecordFoodCreateBulk) SaveX(ctx context.Context) []*RecordFood {
	v, err := rfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rfcb *RecordFoodCreateBulk) Exec(ctx context.Context) error {
	_, err := rfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rfcb *RecordFoodCreateBulk) ExecX(ctx context.Context) {
	if err := rfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
