// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hero/database/ent/predicate"
	"hero/database/ent/user"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetSocialUserID sets the social_user_id field.
func (uu *UserUpdate) SetSocialUserID(s string) *UserUpdate {
	uu.mutation.SetSocialUserID(s)
	return uu
}

// SetNillableSocialUserID sets the social_user_id field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialUserID(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialUserID(*s)
	}
	return uu
}

// SetSocialAvatarURL sets the social_avatar_url field.
func (uu *UserUpdate) SetSocialAvatarURL(s string) *UserUpdate {
	uu.mutation.SetSocialAvatarURL(s)
	return uu
}

// SetNillableSocialAvatarURL sets the social_avatar_url field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialAvatarURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialAvatarURL(*s)
	}
	return uu
}

// SetSocialEmail sets the social_email field.
func (uu *UserUpdate) SetSocialEmail(s string) *UserUpdate {
	uu.mutation.SetSocialEmail(s)
	return uu
}

// SetNillableSocialEmail sets the social_email field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialEmail(*s)
	}
	return uu
}

// SetSocialName sets the social_name field.
func (uu *UserUpdate) SetSocialName(s string) *UserUpdate {
	uu.mutation.SetSocialName(s)
	return uu
}

// SetNillableSocialName sets the social_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialName(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialName(*s)
	}
	return uu
}

// SetSocialType sets the social_type field.
func (uu *UserUpdate) SetSocialType(s string) *UserUpdate {
	uu.mutation.SetSocialType(s)
	return uu
}

// SetNillableSocialType sets the social_type field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialType(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialType(*s)
	}
	return uu
}

// SetSocialPayload sets the social_payload field.
func (uu *UserUpdate) SetSocialPayload(s string) *UserUpdate {
	uu.mutation.SetSocialPayload(s)
	return uu
}

// SetNillableSocialPayload sets the social_payload field if the given value is not nil.
func (uu *UserUpdate) SetNillableSocialPayload(s *string) *UserUpdate {
	if s != nil {
		uu.SetSocialPayload(*s)
	}
	return uu
}

// SetHeroRepeat sets the hero_repeat field.
func (uu *UserUpdate) SetHeroRepeat(u uint) *UserUpdate {
	uu.mutation.ResetHeroRepeat()
	uu.mutation.SetHeroRepeat(u)
	return uu
}

// SetNillableHeroRepeat sets the hero_repeat field if the given value is not nil.
func (uu *UserUpdate) SetNillableHeroRepeat(u *uint) *UserUpdate {
	if u != nil {
		uu.SetHeroRepeat(*u)
	}
	return uu
}

// AddHeroRepeat adds u to hero_repeat.
func (uu *UserUpdate) AddHeroRepeat(u uint) *UserUpdate {
	uu.mutation.AddHeroRepeat(u)
	return uu
}

// SetCreatedAt sets the created_at field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// ClearCreatedAt clears the value of created_at.
func (uu *UserUpdate) ClearCreatedAt() *UserUpdate {
	uu.mutation.ClearCreatedAt()
	return uu
}

// SetUpdatedAt sets the updated_at field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// ClearUpdatedAt clears the value of updated_at.
func (uu *UserUpdate) ClearUpdatedAt() *UserUpdate {
	uu.mutation.ClearUpdatedAt()
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.SocialUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialUserID,
		})
	}
	if value, ok := uu.mutation.SocialAvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialAvatarURL,
		})
	}
	if value, ok := uu.mutation.SocialEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialEmail,
		})
	}
	if value, ok := uu.mutation.SocialName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialName,
		})
	}
	if value, ok := uu.mutation.SocialType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialType,
		})
	}
	if value, ok := uu.mutation.SocialPayload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialPayload,
		})
	}
	if value, ok := uu.mutation.HeroRepeat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldHeroRepeat,
		})
	}
	if value, ok := uu.mutation.AddedHeroRepeat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldHeroRepeat,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if uu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetSocialUserID sets the social_user_id field.
func (uuo *UserUpdateOne) SetSocialUserID(s string) *UserUpdateOne {
	uuo.mutation.SetSocialUserID(s)
	return uuo
}

// SetNillableSocialUserID sets the social_user_id field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialUserID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialUserID(*s)
	}
	return uuo
}

// SetSocialAvatarURL sets the social_avatar_url field.
func (uuo *UserUpdateOne) SetSocialAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetSocialAvatarURL(s)
	return uuo
}

// SetNillableSocialAvatarURL sets the social_avatar_url field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialAvatarURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialAvatarURL(*s)
	}
	return uuo
}

// SetSocialEmail sets the social_email field.
func (uuo *UserUpdateOne) SetSocialEmail(s string) *UserUpdateOne {
	uuo.mutation.SetSocialEmail(s)
	return uuo
}

// SetNillableSocialEmail sets the social_email field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialEmail(*s)
	}
	return uuo
}

// SetSocialName sets the social_name field.
func (uuo *UserUpdateOne) SetSocialName(s string) *UserUpdateOne {
	uuo.mutation.SetSocialName(s)
	return uuo
}

// SetNillableSocialName sets the social_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialName(*s)
	}
	return uuo
}

// SetSocialType sets the social_type field.
func (uuo *UserUpdateOne) SetSocialType(s string) *UserUpdateOne {
	uuo.mutation.SetSocialType(s)
	return uuo
}

// SetNillableSocialType sets the social_type field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialType(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialType(*s)
	}
	return uuo
}

// SetSocialPayload sets the social_payload field.
func (uuo *UserUpdateOne) SetSocialPayload(s string) *UserUpdateOne {
	uuo.mutation.SetSocialPayload(s)
	return uuo
}

// SetNillableSocialPayload sets the social_payload field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSocialPayload(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSocialPayload(*s)
	}
	return uuo
}

// SetHeroRepeat sets the hero_repeat field.
func (uuo *UserUpdateOne) SetHeroRepeat(u uint) *UserUpdateOne {
	uuo.mutation.ResetHeroRepeat()
	uuo.mutation.SetHeroRepeat(u)
	return uuo
}

// SetNillableHeroRepeat sets the hero_repeat field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHeroRepeat(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetHeroRepeat(*u)
	}
	return uuo
}

// AddHeroRepeat adds u to hero_repeat.
func (uuo *UserUpdateOne) AddHeroRepeat(u uint) *UserUpdateOne {
	uuo.mutation.AddHeroRepeat(u)
	return uuo
}

// SetCreatedAt sets the created_at field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// ClearCreatedAt clears the value of created_at.
func (uuo *UserUpdateOne) ClearCreatedAt() *UserUpdateOne {
	uuo.mutation.ClearCreatedAt()
	return uuo
}

// SetUpdatedAt sets the updated_at field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// ClearUpdatedAt clears the value of updated_at.
func (uuo *UserUpdateOne) ClearUpdatedAt() *UserUpdateOne {
	uuo.mutation.ClearUpdatedAt()
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing User.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.SocialUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialUserID,
		})
	}
	if value, ok := uuo.mutation.SocialAvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialAvatarURL,
		})
	}
	if value, ok := uuo.mutation.SocialEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialEmail,
		})
	}
	if value, ok := uuo.mutation.SocialName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialName,
		})
	}
	if value, ok := uuo.mutation.SocialType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialType,
		})
	}
	if value, ok := uuo.mutation.SocialPayload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSocialPayload,
		})
	}
	if value, ok := uuo.mutation.HeroRepeat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldHeroRepeat,
		})
	}
	if value, ok := uuo.mutation.AddedHeroRepeat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: user.FieldHeroRepeat,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedAt,
		})
	}
	if uuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldCreatedAt,
		})
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if uuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldUpdatedAt,
		})
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
