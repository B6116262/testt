// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/prefix"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PersonalID holds the value of the "personalID" field.
	PersonalID int `json:"personalID,omitempty"`
	// PatientName holds the value of the "patientName" field.
	PatientName string `json:"patientName,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// HospitalNumber holds the value of the "hospitalNumber" field.
	HospitalNumber string `json:"hospitalNumber,omitempty"`
	// DrugAllergy holds the value of the "drugAllergy" field.
	DrugAllergy string `json:"drugAllergy,omitempty"`
	// AddedTime holds the value of the "added_time" field.
	AddedTime time.Time `json:"added_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges                            PatientEdges `json:"edges"`
	blood_type_blood_type_to_patient *int
	gender_gender_to_patient         *int
	prefix_prefix_to_patient         *int
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Prefix holds the value of the Prefix edge.
	Prefix *Prefix
	// Gender holds the value of the Gender edge.
	Gender *Gender
	// Bloodtype holds the value of the Bloodtype edge.
	Bloodtype *BloodType
	// PatientToTriageResult holds the value of the patientToTriageResult edge.
	PatientToTriageResult []*TriageResult
	// PatientToAppointmentResults holds the value of the PatientToAppointmentResults edge.
	PatientToAppointmentResults []*AppointmentResults
	// PatientToMedicalProcedure holds the value of the PatientToMedicalProcedure edge.
	PatientToMedicalProcedure []*MedicalProcedure
	// PatientToRightToTreatment holds the value of the PatientToRightToTreatment edge.
	PatientToRightToTreatment []*RightToTreatment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// PrefixOrErr returns the Prefix value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) PrefixOrErr() (*Prefix, error) {
	if e.loadedTypes[0] {
		if e.Prefix == nil {
			// The edge Prefix was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: prefix.Label}
		}
		return e.Prefix, nil
	}
	return nil, &NotLoadedError{edge: "Prefix"}
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[1] {
		if e.Gender == nil {
			// The edge Gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "Gender"}
}

// BloodtypeOrErr returns the Bloodtype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) BloodtypeOrErr() (*BloodType, error) {
	if e.loadedTypes[2] {
		if e.Bloodtype == nil {
			// The edge Bloodtype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bloodtype.Label}
		}
		return e.Bloodtype, nil
	}
	return nil, &NotLoadedError{edge: "Bloodtype"}
}

// PatientToTriageResultOrErr returns the PatientToTriageResult value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) PatientToTriageResultOrErr() ([]*TriageResult, error) {
	if e.loadedTypes[3] {
		return e.PatientToTriageResult, nil
	}
	return nil, &NotLoadedError{edge: "patientToTriageResult"}
}

// PatientToAppointmentResultsOrErr returns the PatientToAppointmentResults value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) PatientToAppointmentResultsOrErr() ([]*AppointmentResults, error) {
	if e.loadedTypes[4] {
		return e.PatientToAppointmentResults, nil
	}
	return nil, &NotLoadedError{edge: "PatientToAppointmentResults"}
}

// PatientToMedicalProcedureOrErr returns the PatientToMedicalProcedure value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) PatientToMedicalProcedureOrErr() ([]*MedicalProcedure, error) {
	if e.loadedTypes[5] {
		return e.PatientToMedicalProcedure, nil
	}
	return nil, &NotLoadedError{edge: "PatientToMedicalProcedure"}
}

// PatientToRightToTreatmentOrErr returns the PatientToRightToTreatment value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) PatientToRightToTreatmentOrErr() ([]*RightToTreatment, error) {
	if e.loadedTypes[6] {
		return e.PatientToRightToTreatment, nil
	}
	return nil, &NotLoadedError{edge: "PatientToRightToTreatment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case patient.FieldID, patient.FieldPersonalID, patient.FieldAge:
			values[i] = &sql.NullInt64{}
		case patient.FieldPatientName, patient.FieldHospitalNumber, patient.FieldDrugAllergy:
			values[i] = &sql.NullString{}
		case patient.FieldAddedTime:
			values[i] = &sql.NullTime{}
		case patient.ForeignKeys[0]: // blood_type_blood_type_to_patient
			values[i] = &sql.NullInt64{}
		case patient.ForeignKeys[1]: // gender_gender_to_patient
			values[i] = &sql.NullInt64{}
		case patient.ForeignKeys[2]: // prefix_prefix_to_patient
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Patient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case patient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case patient.FieldPersonalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field personalID", values[i])
			} else if value.Valid {
				pa.PersonalID = int(value.Int64)
			}
		case patient.FieldPatientName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field patientName", values[i])
			} else if value.Valid {
				pa.PatientName = value.String
			}
		case patient.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pa.Age = int(value.Int64)
			}
		case patient.FieldHospitalNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hospitalNumber", values[i])
			} else if value.Valid {
				pa.HospitalNumber = value.String
			}
		case patient.FieldDrugAllergy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drugAllergy", values[i])
			} else if value.Valid {
				pa.DrugAllergy = value.String
			}
		case patient.FieldAddedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field added_time", values[i])
			} else if value.Valid {
				pa.AddedTime = value.Time
			}
		case patient.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blood_type_blood_type_to_patient", value)
			} else if value.Valid {
				pa.blood_type_blood_type_to_patient = new(int)
				*pa.blood_type_blood_type_to_patient = int(value.Int64)
			}
		case patient.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field gender_gender_to_patient", value)
			} else if value.Valid {
				pa.gender_gender_to_patient = new(int)
				*pa.gender_gender_to_patient = int(value.Int64)
			}
		case patient.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field prefix_prefix_to_patient", value)
			} else if value.Valid {
				pa.prefix_prefix_to_patient = new(int)
				*pa.prefix_prefix_to_patient = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPrefix queries the "Prefix" edge of the Patient entity.
func (pa *Patient) QueryPrefix() *PrefixQuery {
	return (&PatientClient{config: pa.config}).QueryPrefix(pa)
}

// QueryGender queries the "Gender" edge of the Patient entity.
func (pa *Patient) QueryGender() *GenderQuery {
	return (&PatientClient{config: pa.config}).QueryGender(pa)
}

// QueryBloodtype queries the "Bloodtype" edge of the Patient entity.
func (pa *Patient) QueryBloodtype() *BloodTypeQuery {
	return (&PatientClient{config: pa.config}).QueryBloodtype(pa)
}

// QueryPatientToTriageResult queries the "patientToTriageResult" edge of the Patient entity.
func (pa *Patient) QueryPatientToTriageResult() *TriageResultQuery {
	return (&PatientClient{config: pa.config}).QueryPatientToTriageResult(pa)
}

// QueryPatientToAppointmentResults queries the "PatientToAppointmentResults" edge of the Patient entity.
func (pa *Patient) QueryPatientToAppointmentResults() *AppointmentResultsQuery {
	return (&PatientClient{config: pa.config}).QueryPatientToAppointmentResults(pa)
}

// QueryPatientToMedicalProcedure queries the "PatientToMedicalProcedure" edge of the Patient entity.
func (pa *Patient) QueryPatientToMedicalProcedure() *MedicalProcedureQuery {
	return (&PatientClient{config: pa.config}).QueryPatientToMedicalProcedure(pa)
}

// QueryPatientToRightToTreatment queries the "PatientToRightToTreatment" edge of the Patient entity.
func (pa *Patient) QueryPatientToRightToTreatment() *RightToTreatmentQuery {
	return (&PatientClient{config: pa.config}).QueryPatientToRightToTreatment(pa)
}

// Update returns a builder for updating this Patient.
// Note that you need to call Patient.Unwrap() before calling this method if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Patient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", personalID=")
	builder.WriteString(fmt.Sprintf("%v", pa.PersonalID))
	builder.WriteString(", patientName=")
	builder.WriteString(pa.PatientName)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pa.Age))
	builder.WriteString(", hospitalNumber=")
	builder.WriteString(pa.HospitalNumber)
	builder.WriteString(", drugAllergy=")
	builder.WriteString(pa.DrugAllergy)
	builder.WriteString(", added_time=")
	builder.WriteString(pa.AddedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
