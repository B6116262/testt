// Code generated by entc, DO NOT EDIT.

package patient

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PersonalID applies equality check predicate on the "personalID" field. It's identical to PersonalIDEQ.
func PersonalID(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPersonalID), v))
	})
}

// PatientName applies equality check predicate on the "patientName" field. It's identical to PatientNameEQ.
func PatientName(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// HospitalNumber applies equality check predicate on the "hospitalNumber" field. It's identical to HospitalNumberEQ.
func HospitalNumber(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHospitalNumber), v))
	})
}

// DrugAllergy applies equality check predicate on the "drugAllergy" field. It's identical to DrugAllergyEQ.
func DrugAllergy(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugAllergy), v))
	})
}

// AddedTime applies equality check predicate on the "added_time" field. It's identical to AddedTimeEQ.
func AddedTime(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// PersonalIDEQ applies the EQ predicate on the "personalID" field.
func PersonalIDEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPersonalID), v))
	})
}

// PersonalIDNEQ applies the NEQ predicate on the "personalID" field.
func PersonalIDNEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPersonalID), v))
	})
}

// PersonalIDIn applies the In predicate on the "personalID" field.
func PersonalIDIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPersonalID), v...))
	})
}

// PersonalIDNotIn applies the NotIn predicate on the "personalID" field.
func PersonalIDNotIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPersonalID), v...))
	})
}

// PersonalIDGT applies the GT predicate on the "personalID" field.
func PersonalIDGT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPersonalID), v))
	})
}

// PersonalIDGTE applies the GTE predicate on the "personalID" field.
func PersonalIDGTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPersonalID), v))
	})
}

// PersonalIDLT applies the LT predicate on the "personalID" field.
func PersonalIDLT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPersonalID), v))
	})
}

// PersonalIDLTE applies the LTE predicate on the "personalID" field.
func PersonalIDLTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPersonalID), v))
	})
}

// PatientNameEQ applies the EQ predicate on the "patientName" field.
func PatientNameEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// PatientNameNEQ applies the NEQ predicate on the "patientName" field.
func PatientNameNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientName), v))
	})
}

// PatientNameIn applies the In predicate on the "patientName" field.
func PatientNameIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientName), v...))
	})
}

// PatientNameNotIn applies the NotIn predicate on the "patientName" field.
func PatientNameNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientName), v...))
	})
}

// PatientNameGT applies the GT predicate on the "patientName" field.
func PatientNameGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientName), v))
	})
}

// PatientNameGTE applies the GTE predicate on the "patientName" field.
func PatientNameGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientName), v))
	})
}

// PatientNameLT applies the LT predicate on the "patientName" field.
func PatientNameLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientName), v))
	})
}

// PatientNameLTE applies the LTE predicate on the "patientName" field.
func PatientNameLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientName), v))
	})
}

// PatientNameContains applies the Contains predicate on the "patientName" field.
func PatientNameContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPatientName), v))
	})
}

// PatientNameHasPrefix applies the HasPrefix predicate on the "patientName" field.
func PatientNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPatientName), v))
	})
}

// PatientNameHasSuffix applies the HasSuffix predicate on the "patientName" field.
func PatientNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPatientName), v))
	})
}

// PatientNameEqualFold applies the EqualFold predicate on the "patientName" field.
func PatientNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPatientName), v))
	})
}

// PatientNameContainsFold applies the ContainsFold predicate on the "patientName" field.
func PatientNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPatientName), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// HospitalNumberEQ applies the EQ predicate on the "hospitalNumber" field.
func HospitalNumberEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberNEQ applies the NEQ predicate on the "hospitalNumber" field.
func HospitalNumberNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberIn applies the In predicate on the "hospitalNumber" field.
func HospitalNumberIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHospitalNumber), v...))
	})
}

// HospitalNumberNotIn applies the NotIn predicate on the "hospitalNumber" field.
func HospitalNumberNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHospitalNumber), v...))
	})
}

// HospitalNumberGT applies the GT predicate on the "hospitalNumber" field.
func HospitalNumberGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberGTE applies the GTE predicate on the "hospitalNumber" field.
func HospitalNumberGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberLT applies the LT predicate on the "hospitalNumber" field.
func HospitalNumberLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberLTE applies the LTE predicate on the "hospitalNumber" field.
func HospitalNumberLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberContains applies the Contains predicate on the "hospitalNumber" field.
func HospitalNumberContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberHasPrefix applies the HasPrefix predicate on the "hospitalNumber" field.
func HospitalNumberHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberHasSuffix applies the HasSuffix predicate on the "hospitalNumber" field.
func HospitalNumberHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberEqualFold applies the EqualFold predicate on the "hospitalNumber" field.
func HospitalNumberEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHospitalNumber), v))
	})
}

// HospitalNumberContainsFold applies the ContainsFold predicate on the "hospitalNumber" field.
func HospitalNumberContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHospitalNumber), v))
	})
}

// DrugAllergyEQ applies the EQ predicate on the "drugAllergy" field.
func DrugAllergyEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyNEQ applies the NEQ predicate on the "drugAllergy" field.
func DrugAllergyNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyIn applies the In predicate on the "drugAllergy" field.
func DrugAllergyIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugAllergy), v...))
	})
}

// DrugAllergyNotIn applies the NotIn predicate on the "drugAllergy" field.
func DrugAllergyNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugAllergy), v...))
	})
}

// DrugAllergyGT applies the GT predicate on the "drugAllergy" field.
func DrugAllergyGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyGTE applies the GTE predicate on the "drugAllergy" field.
func DrugAllergyGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyLT applies the LT predicate on the "drugAllergy" field.
func DrugAllergyLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyLTE applies the LTE predicate on the "drugAllergy" field.
func DrugAllergyLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyContains applies the Contains predicate on the "drugAllergy" field.
func DrugAllergyContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyHasPrefix applies the HasPrefix predicate on the "drugAllergy" field.
func DrugAllergyHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyHasSuffix applies the HasSuffix predicate on the "drugAllergy" field.
func DrugAllergyHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyEqualFold applies the EqualFold predicate on the "drugAllergy" field.
func DrugAllergyEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugAllergy), v))
	})
}

// DrugAllergyContainsFold applies the ContainsFold predicate on the "drugAllergy" field.
func DrugAllergyContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugAllergy), v))
	})
}

// AddedTimeEQ applies the EQ predicate on the "added_time" field.
func AddedTimeEQ(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeNEQ applies the NEQ predicate on the "added_time" field.
func AddedTimeNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedTime), v))
	})
}

// AddedTimeIn applies the In predicate on the "added_time" field.
func AddedTimeIn(vs ...time.Time) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeNotIn applies the NotIn predicate on the "added_time" field.
func AddedTimeNotIn(vs ...time.Time) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedTime), v...))
	})
}

// AddedTimeGT applies the GT predicate on the "added_time" field.
func AddedTimeGT(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeGTE applies the GTE predicate on the "added_time" field.
func AddedTimeGTE(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLT applies the LT predicate on the "added_time" field.
func AddedTimeLT(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedTime), v))
	})
}

// AddedTimeLTE applies the LTE predicate on the "added_time" field.
func AddedTimeLTE(v time.Time) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedTime), v))
	})
}

// HasPrefix applies the HasEdge predicate on the "Prefix" edge.
func HasPrefix() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefixTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrefixTable, PrefixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrefixWith applies the HasEdge predicate on the "Prefix" edge with a given conditions (other predicates).
func HasPrefixWith(preds ...predicate.Prefix) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrefixTable, PrefixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGender applies the HasEdge predicate on the "Gender" edge.
func HasGender() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderWith applies the HasEdge predicate on the "Gender" edge with a given conditions (other predicates).
func HasGenderWith(preds ...predicate.Gender) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBloodtype applies the HasEdge predicate on the "Bloodtype" edge.
func HasBloodtype() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BloodtypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BloodtypeTable, BloodtypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBloodtypeWith applies the HasEdge predicate on the "Bloodtype" edge with a given conditions (other predicates).
func HasBloodtypeWith(preds ...predicate.BloodType) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BloodtypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BloodtypeTable, BloodtypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientToTriageResult applies the HasEdge predicate on the "patientToTriageResult" edge.
func HasPatientToTriageResult() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToTriageResultTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToTriageResultTable, PatientToTriageResultColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientToTriageResultWith applies the HasEdge predicate on the "patientToTriageResult" edge with a given conditions (other predicates).
func HasPatientToTriageResultWith(preds ...predicate.TriageResult) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToTriageResultInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToTriageResultTable, PatientToTriageResultColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientToAppointmentResults applies the HasEdge predicate on the "PatientToAppointmentResults" edge.
func HasPatientToAppointmentResults() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToAppointmentResultsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToAppointmentResultsTable, PatientToAppointmentResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientToAppointmentResultsWith applies the HasEdge predicate on the "PatientToAppointmentResults" edge with a given conditions (other predicates).
func HasPatientToAppointmentResultsWith(preds ...predicate.AppointmentResults) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToAppointmentResultsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToAppointmentResultsTable, PatientToAppointmentResultsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientToMedicalProcedure applies the HasEdge predicate on the "PatientToMedicalProcedure" edge.
func HasPatientToMedicalProcedure() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToMedicalProcedureTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToMedicalProcedureTable, PatientToMedicalProcedureColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientToMedicalProcedureWith applies the HasEdge predicate on the "PatientToMedicalProcedure" edge with a given conditions (other predicates).
func HasPatientToMedicalProcedureWith(preds ...predicate.MedicalProcedure) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToMedicalProcedureInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToMedicalProcedureTable, PatientToMedicalProcedureColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientToRightToTreatment applies the HasEdge predicate on the "PatientToRightToTreatment" edge.
func HasPatientToRightToTreatment() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToRightToTreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToRightToTreatmentTable, PatientToRightToTreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientToRightToTreatmentWith applies the HasEdge predicate on the "PatientToRightToTreatment" edge with a given conditions (other predicates).
func HasPatientToRightToTreatmentWith(preds ...predicate.RightToTreatment) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientToRightToTreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientToRightToTreatmentTable, PatientToRightToTreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		p(s.Not())
	})
}