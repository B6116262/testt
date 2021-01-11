// Code generated by entc, DO NOT EDIT.

package urgencylevel

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UrgencyName applies equality check predicate on the "urgencyName" field. It's identical to UrgencyNameEQ.
func UrgencyName(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameEQ applies the EQ predicate on the "urgencyName" field.
func UrgencyNameEQ(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameNEQ applies the NEQ predicate on the "urgencyName" field.
func UrgencyNameNEQ(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameIn applies the In predicate on the "urgencyName" field.
func UrgencyNameIn(vs ...string) predicate.UrgencyLevel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUrgencyName), v...))
	})
}

// UrgencyNameNotIn applies the NotIn predicate on the "urgencyName" field.
func UrgencyNameNotIn(vs ...string) predicate.UrgencyLevel {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUrgencyName), v...))
	})
}

// UrgencyNameGT applies the GT predicate on the "urgencyName" field.
func UrgencyNameGT(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameGTE applies the GTE predicate on the "urgencyName" field.
func UrgencyNameGTE(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameLT applies the LT predicate on the "urgencyName" field.
func UrgencyNameLT(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameLTE applies the LTE predicate on the "urgencyName" field.
func UrgencyNameLTE(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameContains applies the Contains predicate on the "urgencyName" field.
func UrgencyNameContains(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameHasPrefix applies the HasPrefix predicate on the "urgencyName" field.
func UrgencyNameHasPrefix(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameHasSuffix applies the HasSuffix predicate on the "urgencyName" field.
func UrgencyNameHasSuffix(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameEqualFold applies the EqualFold predicate on the "urgencyName" field.
func UrgencyNameEqualFold(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUrgencyName), v))
	})
}

// UrgencyNameContainsFold applies the ContainsFold predicate on the "urgencyName" field.
func UrgencyNameContainsFold(v string) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUrgencyName), v))
	})
}

// HasUrgencyLevelToTriageResult applies the HasEdge predicate on the "urgencyLevelToTriageResult" edge.
func HasUrgencyLevelToTriageResult() predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UrgencyLevelToTriageResultTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UrgencyLevelToTriageResultTable, UrgencyLevelToTriageResultColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUrgencyLevelToTriageResultWith applies the HasEdge predicate on the "urgencyLevelToTriageResult" edge with a given conditions (other predicates).
func HasUrgencyLevelToTriageResultWith(preds ...predicate.TriageResult) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UrgencyLevelToTriageResultInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UrgencyLevelToTriageResultTable, UrgencyLevelToTriageResultColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UrgencyLevel) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UrgencyLevel) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
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
func Not(p predicate.UrgencyLevel) predicate.UrgencyLevel {
	return predicate.UrgencyLevel(func(s *sql.Selector) {
		p(s.Not())
	})
}
