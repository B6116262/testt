// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/nurse"
)

// Nurse is the model entity for the Nurse schema.
type Nurse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NurseName holds the value of the "nurseName" field.
	NurseName string `json:"nurseName,omitempty"`
	// NurseUsername holds the value of the "nurseUsername" field.
	NurseUsername string `json:"nurseUsername,omitempty"`
	// NursePassword holds the value of the "nursePassword" field.
	NursePassword string `json:"nursePassword,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NurseQuery when eager-loading is set.
	Edges NurseEdges `json:"edges"`
}

// NurseEdges holds the relations/edges for other nodes in the graph.
type NurseEdges struct {
	// NurseToTriageResult holds the value of the nurseToTriageResult edge.
	NurseToTriageResult []*TriageResult
	// NurseToAppointmentResults holds the value of the NurseToAppointmentResults edge.
	NurseToAppointmentResults []*AppointmentResults
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NurseToTriageResultOrErr returns the NurseToTriageResult value or an error if the edge
// was not loaded in eager-loading.
func (e NurseEdges) NurseToTriageResultOrErr() ([]*TriageResult, error) {
	if e.loadedTypes[0] {
		return e.NurseToTriageResult, nil
	}
	return nil, &NotLoadedError{edge: "nurseToTriageResult"}
}

// NurseToAppointmentResultsOrErr returns the NurseToAppointmentResults value or an error if the edge
// was not loaded in eager-loading.
func (e NurseEdges) NurseToAppointmentResultsOrErr() ([]*AppointmentResults, error) {
	if e.loadedTypes[1] {
		return e.NurseToAppointmentResults, nil
	}
	return nil, &NotLoadedError{edge: "NurseToAppointmentResults"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Nurse) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case nurse.FieldID:
			values[i] = &sql.NullInt64{}
		case nurse.FieldNurseName, nurse.FieldNurseUsername, nurse.FieldNursePassword:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Nurse", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Nurse fields.
func (n *Nurse) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nurse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case nurse.FieldNurseName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nurseName", values[i])
			} else if value.Valid {
				n.NurseName = value.String
			}
		case nurse.FieldNurseUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nurseUsername", values[i])
			} else if value.Valid {
				n.NurseUsername = value.String
			}
		case nurse.FieldNursePassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nursePassword", values[i])
			} else if value.Valid {
				n.NursePassword = value.String
			}
		}
	}
	return nil
}

// QueryNurseToTriageResult queries the "nurseToTriageResult" edge of the Nurse entity.
func (n *Nurse) QueryNurseToTriageResult() *TriageResultQuery {
	return (&NurseClient{config: n.config}).QueryNurseToTriageResult(n)
}

// QueryNurseToAppointmentResults queries the "NurseToAppointmentResults" edge of the Nurse entity.
func (n *Nurse) QueryNurseToAppointmentResults() *AppointmentResultsQuery {
	return (&NurseClient{config: n.config}).QueryNurseToAppointmentResults(n)
}

// Update returns a builder for updating this Nurse.
// Note that you need to call Nurse.Unwrap() before calling this method if this Nurse
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Nurse) Update() *NurseUpdateOne {
	return (&NurseClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Nurse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Nurse) Unwrap() *Nurse {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Nurse is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Nurse) String() string {
	var builder strings.Builder
	builder.WriteString("Nurse(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", nurseName=")
	builder.WriteString(n.NurseName)
	builder.WriteString(", nurseUsername=")
	builder.WriteString(n.NurseUsername)
	builder.WriteString(", nursePassword=")
	builder.WriteString(n.NursePassword)
	builder.WriteByte(')')
	return builder.String()
}

// Nurses is a parsable slice of Nurse.
type Nurses []*Nurse

func (n Nurses) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}