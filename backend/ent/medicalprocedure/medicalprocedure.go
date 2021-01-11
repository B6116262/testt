// Code generated by entc, DO NOT EDIT.

package medicalprocedure

const (
	// Label holds the string label denoting the medicalprocedure type in the database.
	Label = "medical_procedure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddtime holds the string denoting the addtime field in the database.
	FieldAddtime = "addtime"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"
	// EdgeProcedureType holds the string denoting the proceduretype edge name in mutations.
	EdgeProcedureType = "ProcedureType"
	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "Doctor"

	// Table holds the table name of the medicalprocedure in the database.
	Table = "medical_procedures"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "medical_procedures"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "patient_patient_to_medical_procedure"
	// ProcedureTypeTable is the table the holds the ProcedureType relation/edge.
	ProcedureTypeTable = "medical_procedures"
	// ProcedureTypeInverseTable is the table name for the ProcedureType entity.
	// It exists in this package in order to avoid circular dependency with the "proceduretype" package.
	ProcedureTypeInverseTable = "procedure_types"
	// ProcedureTypeColumn is the table column denoting the ProcedureType relation/edge.
	ProcedureTypeColumn = "procedure_type_procedure_to_medical_procedure"
	// DoctorTable is the table the holds the Doctor relation/edge.
	DoctorTable = "medical_procedures"
	// DoctorInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorInverseTable = "doctors"
	// DoctorColumn is the table column denoting the Doctor relation/edge.
	DoctorColumn = "doctor_doctor_to_medical_procedure"
)

// Columns holds all SQL columns for medicalprocedure fields.
var Columns = []string{
	FieldID,
	FieldAddtime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the MedicalProcedure type.
var ForeignKeys = []string{
	"doctor_doctor_to_medical_procedure",
	"patient_patient_to_medical_procedure",
	"procedure_type_procedure_to_medical_procedure",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
