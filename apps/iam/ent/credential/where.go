// Code generated by ent, DO NOT EDIT.

package credential

import (
	"infra-kit/apps/iam/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldID, id))
}

// Ak applies equality check predicate on the "ak" field. It's identical to AkEQ.
func Ak(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldAk, v))
}

// Sk applies equality check predicate on the "sk" field. It's identical to SkEQ.
func Sk(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldSk, v))
}

// Ctime applies equality check predicate on the "ctime" field. It's identical to CtimeEQ.
func Ctime(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCtime, v))
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldMtime, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldDeleted, v))
}

// AkEQ applies the EQ predicate on the "ak" field.
func AkEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldAk, v))
}

// AkNEQ applies the NEQ predicate on the "ak" field.
func AkNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldAk, v))
}

// AkIn applies the In predicate on the "ak" field.
func AkIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldAk, vs...))
}

// AkNotIn applies the NotIn predicate on the "ak" field.
func AkNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldAk, vs...))
}

// AkGT applies the GT predicate on the "ak" field.
func AkGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldAk, v))
}

// AkGTE applies the GTE predicate on the "ak" field.
func AkGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldAk, v))
}

// AkLT applies the LT predicate on the "ak" field.
func AkLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldAk, v))
}

// AkLTE applies the LTE predicate on the "ak" field.
func AkLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldAk, v))
}

// AkContains applies the Contains predicate on the "ak" field.
func AkContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldAk, v))
}

// AkHasPrefix applies the HasPrefix predicate on the "ak" field.
func AkHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldAk, v))
}

// AkHasSuffix applies the HasSuffix predicate on the "ak" field.
func AkHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldAk, v))
}

// AkEqualFold applies the EqualFold predicate on the "ak" field.
func AkEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldAk, v))
}

// AkContainsFold applies the ContainsFold predicate on the "ak" field.
func AkContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldAk, v))
}

// SkEQ applies the EQ predicate on the "sk" field.
func SkEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldSk, v))
}

// SkNEQ applies the NEQ predicate on the "sk" field.
func SkNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldSk, v))
}

// SkIn applies the In predicate on the "sk" field.
func SkIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldSk, vs...))
}

// SkNotIn applies the NotIn predicate on the "sk" field.
func SkNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldSk, vs...))
}

// SkGT applies the GT predicate on the "sk" field.
func SkGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldSk, v))
}

// SkGTE applies the GTE predicate on the "sk" field.
func SkGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldSk, v))
}

// SkLT applies the LT predicate on the "sk" field.
func SkLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldSk, v))
}

// SkLTE applies the LTE predicate on the "sk" field.
func SkLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldSk, v))
}

// SkContains applies the Contains predicate on the "sk" field.
func SkContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldSk, v))
}

// SkHasPrefix applies the HasPrefix predicate on the "sk" field.
func SkHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldSk, v))
}

// SkHasSuffix applies the HasSuffix predicate on the "sk" field.
func SkHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldSk, v))
}

// SkEqualFold applies the EqualFold predicate on the "sk" field.
func SkEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldSk, v))
}

// SkContainsFold applies the ContainsFold predicate on the "sk" field.
func SkContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldSk, v))
}

// CtimeEQ applies the EQ predicate on the "ctime" field.
func CtimeEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCtime, v))
}

// CtimeNEQ applies the NEQ predicate on the "ctime" field.
func CtimeNEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldCtime, v))
}

// CtimeIn applies the In predicate on the "ctime" field.
func CtimeIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldCtime, vs...))
}

// CtimeNotIn applies the NotIn predicate on the "ctime" field.
func CtimeNotIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldCtime, vs...))
}

// CtimeGT applies the GT predicate on the "ctime" field.
func CtimeGT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldCtime, v))
}

// CtimeGTE applies the GTE predicate on the "ctime" field.
func CtimeGTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldCtime, v))
}

// CtimeLT applies the LT predicate on the "ctime" field.
func CtimeLT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldCtime, v))
}

// CtimeLTE applies the LTE predicate on the "ctime" field.
func CtimeLTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldCtime, v))
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldMtime, v))
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldMtime, v))
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldMtime, vs...))
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldMtime, vs...))
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldMtime, v))
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldMtime, v))
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldMtime, v))
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldMtime, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldDeleted, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Credential) predicate.Credential {
	return predicate.Credential(sql.NotPredicates(p))
}
