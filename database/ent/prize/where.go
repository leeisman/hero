// Code generated by entc, DO NOT EDIT.

package prize

import (
	"hero/database/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SocialUserID applies equality check predicate on the "social_user_id" field. It's identical to SocialUserIDEQ.
func SocialUserID(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSocialUserID), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// SocialUserIDEQ applies the EQ predicate on the "social_user_id" field.
func SocialUserIDEQ(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDNEQ applies the NEQ predicate on the "social_user_id" field.
func SocialUserIDNEQ(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDIn applies the In predicate on the "social_user_id" field.
func SocialUserIDIn(vs ...string) predicate.Prize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSocialUserID), v...))
	})
}

// SocialUserIDNotIn applies the NotIn predicate on the "social_user_id" field.
func SocialUserIDNotIn(vs ...string) predicate.Prize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSocialUserID), v...))
	})
}

// SocialUserIDGT applies the GT predicate on the "social_user_id" field.
func SocialUserIDGT(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDGTE applies the GTE predicate on the "social_user_id" field.
func SocialUserIDGTE(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDLT applies the LT predicate on the "social_user_id" field.
func SocialUserIDLT(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDLTE applies the LTE predicate on the "social_user_id" field.
func SocialUserIDLTE(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDContains applies the Contains predicate on the "social_user_id" field.
func SocialUserIDContains(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDHasPrefix applies the HasPrefix predicate on the "social_user_id" field.
func SocialUserIDHasPrefix(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDHasSuffix applies the HasSuffix predicate on the "social_user_id" field.
func SocialUserIDHasSuffix(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDEqualFold applies the EqualFold predicate on the "social_user_id" field.
func SocialUserIDEqualFold(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSocialUserID), v))
	})
}

// SocialUserIDContainsFold applies the ContainsFold predicate on the "social_user_id" field.
func SocialUserIDContainsFold(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSocialUserID), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Prize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Prize {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prize(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDate), v))
	})
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDate), v))
	})
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDate), v))
	})
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDate)))
	})
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDate)))
	})
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDate), v))
	})
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDate), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Prize) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Prize) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
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
func Not(p predicate.Prize) predicate.Prize {
	return predicate.Prize(func(s *sql.Selector) {
		p(s.Not())
	})
}
