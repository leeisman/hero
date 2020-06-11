// Code generated by entc, DO NOT EDIT.

package useractiverecord

import (
	"hero/database/ent/predicate"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ActiveType applies equality check predicate on the "active_type" field. It's identical to ActiveTypeEQ.
func ActiveType(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActiveType), v))
	})
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// IsFinished applies equality check predicate on the "is_finished" field. It's identical to IsFinishedEQ.
func IsFinished(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFinished), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// ActiveTypeEQ applies the EQ predicate on the "active_type" field.
func ActiveTypeEQ(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActiveType), v))
	})
}

// ActiveTypeNEQ applies the NEQ predicate on the "active_type" field.
func ActiveTypeNEQ(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActiveType), v))
	})
}

// ActiveTypeIn applies the In predicate on the "active_type" field.
func ActiveTypeIn(vs ...string) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActiveType), v...))
	})
}

// ActiveTypeNotIn applies the NotIn predicate on the "active_type" field.
func ActiveTypeNotIn(vs ...string) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActiveType), v...))
	})
}

// ActiveTypeGT applies the GT predicate on the "active_type" field.
func ActiveTypeGT(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActiveType), v))
	})
}

// ActiveTypeGTE applies the GTE predicate on the "active_type" field.
func ActiveTypeGTE(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActiveType), v))
	})
}

// ActiveTypeLT applies the LT predicate on the "active_type" field.
func ActiveTypeLT(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActiveType), v))
	})
}

// ActiveTypeLTE applies the LTE predicate on the "active_type" field.
func ActiveTypeLTE(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActiveType), v))
	})
}

// ActiveTypeContains applies the Contains predicate on the "active_type" field.
func ActiveTypeContains(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActiveType), v))
	})
}

// ActiveTypeHasPrefix applies the HasPrefix predicate on the "active_type" field.
func ActiveTypeHasPrefix(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActiveType), v))
	})
}

// ActiveTypeHasSuffix applies the HasSuffix predicate on the "active_type" field.
func ActiveTypeHasSuffix(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActiveType), v))
	})
}

// ActiveTypeEqualFold applies the EqualFold predicate on the "active_type" field.
func ActiveTypeEqualFold(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActiveType), v))
	})
}

// ActiveTypeContainsFold applies the ContainsFold predicate on the "active_type" field.
func ActiveTypeContainsFold(v string) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActiveType), v))
	})
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScore), v))
	})
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScore), v...))
	})
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScore), v...))
	})
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScore), v))
	})
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScore), v))
	})
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScore), v))
	})
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScore), v))
	})
}

// IsFinishedEQ applies the EQ predicate on the "is_finished" field.
func IsFinishedEQ(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsFinished), v))
	})
}

// IsFinishedNEQ applies the NEQ predicate on the "is_finished" field.
func IsFinishedNEQ(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsFinished), v))
	})
}

// IsFinishedIn applies the In predicate on the "is_finished" field.
func IsFinishedIn(vs ...uint) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIsFinished), v...))
	})
}

// IsFinishedNotIn applies the NotIn predicate on the "is_finished" field.
func IsFinishedNotIn(vs ...uint) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIsFinished), v...))
	})
}

// IsFinishedGT applies the GT predicate on the "is_finished" field.
func IsFinishedGT(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsFinished), v))
	})
}

// IsFinishedGTE applies the GTE predicate on the "is_finished" field.
func IsFinishedGTE(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsFinished), v))
	})
}

// IsFinishedLT applies the LT predicate on the "is_finished" field.
func IsFinishedLT(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsFinished), v))
	})
}

// IsFinishedLTE applies the LTE predicate on the "is_finished" field.
func IsFinishedLTE(v uint) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsFinished), v))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndedAt), v...))
	})
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndedAt), v...))
	})
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndedAt), v))
	})
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndedAt), v))
	})
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtIsNil applies the IsNil predicate on the "ended_at" field.
func EndedAtIsNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndedAt)))
	})
}

// EndedAtNotNil applies the NotNil predicate on the "ended_at" field.
func EndedAtNotNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserActiveRecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.UserActiveRecord) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.UserActiveRecord) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserActiveRecord) predicate.UserActiveRecord {
	return predicate.UserActiveRecord(func(s *sql.Selector) {
		p(s.Not())
	})
}
