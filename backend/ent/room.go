// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/room"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoomName holds the value of the "roomName" field.
	RoomName string `json:"roomName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// RoomToAppointmentResults holds the value of the RoomToAppointmentResults edge.
	RoomToAppointmentResults []*AppointmentResults
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomToAppointmentResultsOrErr returns the RoomToAppointmentResults value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) RoomToAppointmentResultsOrErr() ([]*AppointmentResults, error) {
	if e.loadedTypes[0] {
		return e.RoomToAppointmentResults, nil
	}
	return nil, &NotLoadedError{edge: "RoomToAppointmentResults"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			values[i] = &sql.NullInt64{}
		case room.FieldRoomName:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Room", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case room.FieldRoomName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field roomName", values[i])
			} else if value.Valid {
				r.RoomName = value.String
			}
		}
	}
	return nil
}

// QueryRoomToAppointmentResults queries the "RoomToAppointmentResults" edge of the Room entity.
func (r *Room) QueryRoomToAppointmentResults() *AppointmentResultsQuery {
	return (&RoomClient{config: r.config}).QueryRoomToAppointmentResults(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", roomName=")
	builder.WriteString(r.RoomName)
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
