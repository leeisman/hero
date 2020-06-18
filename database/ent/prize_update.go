// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hero/database/ent/predicate"
	"hero/database/ent/prize"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PrizeUpdate is the builder for updating Prize entities.
type PrizeUpdate struct {
	config
	hooks      []Hook
	mutation   *PrizeMutation
	predicates []predicate.Prize
}

// Where adds a new predicate for the builder.
func (pu *PrizeUpdate) Where(ps ...predicate.Prize) *PrizeUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetSocialUserID sets the social_user_id field.
func (pu *PrizeUpdate) SetSocialUserID(s string) *PrizeUpdate {
	pu.mutation.SetSocialUserID(s)
	return pu
}

// SetNillableSocialUserID sets the social_user_id field if the given value is not nil.
func (pu *PrizeUpdate) SetNillableSocialUserID(s *string) *PrizeUpdate {
	if s != nil {
		pu.SetSocialUserID(*s)
	}
	return pu
}

// SetDate sets the date field.
func (pu *PrizeUpdate) SetDate(s string) *PrizeUpdate {
	pu.mutation.SetDate(s)
	return pu
}

// SetNillableDate sets the date field if the given value is not nil.
func (pu *PrizeUpdate) SetNillableDate(s *string) *PrizeUpdate {
	if s != nil {
		pu.SetDate(*s)
	}
	return pu
}

// ClearDate clears the value of date.
func (pu *PrizeUpdate) ClearDate() *PrizeUpdate {
	pu.mutation.ClearDate()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrizeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrizeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrizeUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrizeUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrizeUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrizeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prize.Table,
			Columns: prize.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prize.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.SocialUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prize.FieldSocialUserID,
		})
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prize.FieldDate,
		})
	}
	if pu.mutation.DateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: prize.FieldDate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prize.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrizeUpdateOne is the builder for updating a single Prize entity.
type PrizeUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrizeMutation
}

// SetSocialUserID sets the social_user_id field.
func (puo *PrizeUpdateOne) SetSocialUserID(s string) *PrizeUpdateOne {
	puo.mutation.SetSocialUserID(s)
	return puo
}

// SetNillableSocialUserID sets the social_user_id field if the given value is not nil.
func (puo *PrizeUpdateOne) SetNillableSocialUserID(s *string) *PrizeUpdateOne {
	if s != nil {
		puo.SetSocialUserID(*s)
	}
	return puo
}

// SetDate sets the date field.
func (puo *PrizeUpdateOne) SetDate(s string) *PrizeUpdateOne {
	puo.mutation.SetDate(s)
	return puo
}

// SetNillableDate sets the date field if the given value is not nil.
func (puo *PrizeUpdateOne) SetNillableDate(s *string) *PrizeUpdateOne {
	if s != nil {
		puo.SetDate(*s)
	}
	return puo
}

// ClearDate clears the value of date.
func (puo *PrizeUpdateOne) ClearDate() *PrizeUpdateOne {
	puo.mutation.ClearDate()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PrizeUpdateOne) Save(ctx context.Context) (*Prize, error) {
	var (
		err  error
		node *Prize
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrizeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrizeUpdateOne) SaveX(ctx context.Context) *Prize {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrizeUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrizeUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrizeUpdateOne) sqlSave(ctx context.Context) (pr *Prize, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prize.Table,
			Columns: prize.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prize.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Prize.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.SocialUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prize.FieldSocialUserID,
		})
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prize.FieldDate,
		})
	}
	if puo.mutation.DateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: prize.FieldDate,
		})
	}
	pr = &Prize{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prize.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
