package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Diagnosis holds the schema definition for the Diagnosis entity.
type Diagnosis struct {
	ent.Schema
}

// Fields of the Diagnosis.
func (Diagnosis) Fields() []ent.Field {
	return nil
}

// Edges of the Diagnosis.
func (Diagnosis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("diagnosis", Doctor.Type).Ref("DoctorToDiagnosis"),
	}
}