// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AppointmentResultsColumns holds the columns for the "appointment_results" table.
	AppointmentResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addtime_appoint", Type: field.TypeTime},
		{Name: "addtime_save", Type: field.TypeTime},
		{Name: "doctor_doctor_to_appointment_results", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_nurse_to_appointment_results", Type: field.TypeInt, Nullable: true},
		{Name: "patient_patient_to_appointment_results", Type: field.TypeInt, Nullable: true},
		{Name: "room_room_to_appointment_results", Type: field.TypeInt, Nullable: true},
	}
	// AppointmentResultsTable holds the schema information for the "appointment_results" table.
	AppointmentResultsTable = &schema.Table{
		Name:       "appointment_results",
		Columns:    AppointmentResultsColumns,
		PrimaryKey: []*schema.Column{AppointmentResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "appointment_results_doctors_DoctorToAppointmentResults",
				Columns: []*schema.Column{AppointmentResultsColumns[3]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointment_results_nurses_NurseToAppointmentResults",
				Columns: []*schema.Column{AppointmentResultsColumns[4]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointment_results_patients_PatientToAppointmentResults",
				Columns: []*schema.Column{AppointmentResultsColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "appointment_results_rooms_RoomToAppointmentResults",
				Columns: []*schema.Column{AppointmentResultsColumns[6]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BloodTypesColumns holds the columns for the "blood_types" table.
	BloodTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "blood", Type: field.TypeString},
	}
	// BloodTypesTable holds the schema information for the "blood_types" table.
	BloodTypesTable = &schema.Table{
		Name:        "blood_types",
		Columns:     BloodTypesColumns,
		PrimaryKey:  []*schema.Column{BloodTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department_name", Type: field.TypeString, Unique: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DiagnosesColumns holds the columns for the "diagnoses" table.
	DiagnosesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DiagnosesTable holds the schema information for the "diagnoses" table.
	DiagnosesTable = &schema.Table{
		Name:        "diagnoses",
		Columns:     DiagnosesColumns,
		PrimaryKey:  []*schema.Column{DiagnosesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorsColumns holds the columns for the "doctors" table.
	DoctorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doctor_name", Type: field.TypeString},
		{Name: "doctor_username", Type: field.TypeString, Unique: true},
		{Name: "doctor_password", Type: field.TypeString},
	}
	// DoctorsTable holds the schema information for the "doctors" table.
	DoctorsTable = &schema.Table{
		Name:        "doctors",
		Columns:     DoctorsColumns,
		PrimaryKey:  []*schema.Column{DoctorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HospitalsColumns holds the columns for the "hospitals" table.
	HospitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hospital_name", Type: field.TypeString, Unique: true},
	}
	// HospitalsTable holds the schema information for the "hospitals" table.
	HospitalsTable = &schema.Table{
		Name:        "hospitals",
		Columns:     HospitalsColumns,
		PrimaryKey:  []*schema.Column{HospitalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalProceduresColumns holds the columns for the "medical_procedures" table.
	MedicalProceduresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addtime", Type: field.TypeTime},
		{Name: "doctor_doctor_to_medical_procedure", Type: field.TypeInt, Nullable: true},
		{Name: "patient_patient_to_medical_procedure", Type: field.TypeInt, Nullable: true},
		{Name: "procedure_type_procedure_to_medical_procedure", Type: field.TypeInt, Nullable: true},
	}
	// MedicalProceduresTable holds the schema information for the "medical_procedures" table.
	MedicalProceduresTable = &schema.Table{
		Name:       "medical_procedures",
		Columns:    MedicalProceduresColumns,
		PrimaryKey: []*schema.Column{MedicalProceduresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "medical_procedures_doctors_DoctorToMedicalProcedure",
				Columns: []*schema.Column{MedicalProceduresColumns[2]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medical_procedures_patients_PatientToMedicalProcedure",
				Columns: []*schema.Column{MedicalProceduresColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medical_procedures_procedure_types_ProcedureToMedicalProcedure",
				Columns: []*schema.Column{MedicalProceduresColumns[4]},

				RefColumns: []*schema.Column{ProcedureTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MedicalRecordsColumns holds the columns for the "medical_records" table.
	MedicalRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// MedicalRecordsTable holds the schema information for the "medical_records" table.
	MedicalRecordsTable = &schema.Table{
		Name:        "medical_records",
		Columns:     MedicalRecordsColumns,
		PrimaryKey:  []*schema.Column{MedicalRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NursesColumns holds the columns for the "nurses" table.
	NursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nurse_name", Type: field.TypeString},
		{Name: "nurse_username", Type: field.TypeString, Unique: true},
		{Name: "nurse_password", Type: field.TypeString},
	}
	// NursesTable holds the schema information for the "nurses" table.
	NursesTable = &schema.Table{
		Name:        "nurses",
		Columns:     NursesColumns,
		PrimaryKey:  []*schema.Column{NursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "personal_id", Type: field.TypeInt, Unique: true},
		{Name: "patient_name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "hospital_number", Type: field.TypeString, Unique: true},
		{Name: "drug_allergy", Type: field.TypeString},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "blood_type_blood_type_to_patient", Type: field.TypeInt, Nullable: true},
		{Name: "gender_gender_to_patient", Type: field.TypeInt, Nullable: true},
		{Name: "prefix_prefix_to_patient", Type: field.TypeInt, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patients_blood_types_BloodTypeToPatient",
				Columns: []*schema.Column{PatientsColumns[7]},

				RefColumns: []*schema.Column{BloodTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_genders_GenderToPatient",
				Columns: []*schema.Column{PatientsColumns[8]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_prefixes_PrefixToPatient",
				Columns: []*schema.Column{PatientsColumns[9]},

				RefColumns: []*schema.Column{PrefixesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrefixesColumns holds the columns for the "prefixes" table.
	PrefixesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "prefix", Type: field.TypeString},
	}
	// PrefixesTable holds the schema information for the "prefixes" table.
	PrefixesTable = &schema.Table{
		Name:        "prefixes",
		Columns:     PrefixesColumns,
		PrimaryKey:  []*schema.Column{PrefixesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProcedureTypesColumns holds the columns for the "procedure_types" table.
	ProcedureTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "procedure_type", Type: field.TypeString},
	}
	// ProcedureTypesTable holds the schema information for the "procedure_types" table.
	ProcedureTypesTable = &schema.Table{
		Name:        "procedure_types",
		Columns:     ProcedureTypesColumns,
		PrimaryKey:  []*schema.Column{ProcedureTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RightToTreatmentsColumns holds the columns for the "right_to_treatments" table.
	RightToTreatmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "starttime", Type: field.TypeTime},
		{Name: "endtime", Type: field.TypeTime},
		{Name: "hospital_hospital", Type: field.TypeInt, Nullable: true},
		{Name: "patient_patient_to_right_to_treatment", Type: field.TypeInt, Nullable: true},
		{Name: "right_to_treatment_type_type", Type: field.TypeInt, Nullable: true},
	}
	// RightToTreatmentsTable holds the schema information for the "right_to_treatments" table.
	RightToTreatmentsTable = &schema.Table{
		Name:       "right_to_treatments",
		Columns:    RightToTreatmentsColumns,
		PrimaryKey: []*schema.Column{RightToTreatmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "right_to_treatments_hospitals_hospital",
				Columns: []*schema.Column{RightToTreatmentsColumns[3]},

				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "right_to_treatments_patients_PatientToRightToTreatment",
				Columns: []*schema.Column{RightToTreatmentsColumns[4]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "right_to_treatments_right_to_treatment_types_type",
				Columns: []*schema.Column{RightToTreatmentsColumns[5]},

				RefColumns: []*schema.Column{RightToTreatmentTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RightToTreatmentTypesColumns holds the columns for the "right_to_treatment_types" table.
	RightToTreatmentTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type_name", Type: field.TypeString, Unique: true},
	}
	// RightToTreatmentTypesTable holds the schema information for the "right_to_treatment_types" table.
	RightToTreatmentTypesTable = &schema.Table{
		Name:        "right_to_treatment_types",
		Columns:     RightToTreatmentTypesColumns,
		PrimaryKey:  []*schema.Column{RightToTreatmentTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "room_name", Type: field.TypeString},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:        "rooms",
		Columns:     RoomsColumns,
		PrimaryKey:  []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TriageResultsColumns holds the columns for the "triage_results" table.
	TriageResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symptom", Type: field.TypeString},
		{Name: "triage_date", Type: field.TypeTime},
		{Name: "department_department_to_triage_result", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_nurse_to_triage_result", Type: field.TypeInt, Nullable: true},
		{Name: "patient_patient_to_triage_result", Type: field.TypeInt, Nullable: true},
		{Name: "urgency_level_urgency_level_to_triage_result", Type: field.TypeInt, Nullable: true},
	}
	// TriageResultsTable holds the schema information for the "triage_results" table.
	TriageResultsTable = &schema.Table{
		Name:       "triage_results",
		Columns:    TriageResultsColumns,
		PrimaryKey: []*schema.Column{TriageResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "triage_results_departments_departmentToTriageResult",
				Columns: []*schema.Column{TriageResultsColumns[3]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "triage_results_nurses_nurseToTriageResult",
				Columns: []*schema.Column{TriageResultsColumns[4]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "triage_results_patients_patientToTriageResult",
				Columns: []*schema.Column{TriageResultsColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "triage_results_urgency_levels_urgencyLevelToTriageResult",
				Columns: []*schema.Column{TriageResultsColumns[6]},

				RefColumns: []*schema.Column{UrgencyLevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UrgencyLevelsColumns holds the columns for the "urgency_levels" table.
	UrgencyLevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "urgency_name", Type: field.TypeString, Unique: true},
	}
	// UrgencyLevelsTable holds the schema information for the "urgency_levels" table.
	UrgencyLevelsTable = &schema.Table{
		Name:        "urgency_levels",
		Columns:     UrgencyLevelsColumns,
		PrimaryKey:  []*schema.Column{UrgencyLevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorDoctorToDiagnosisColumns holds the columns for the "doctor_DoctorToDiagnosis" table.
	DoctorDoctorToDiagnosisColumns = []*schema.Column{
		{Name: "doctor_id", Type: field.TypeInt},
		{Name: "diagnosis_id", Type: field.TypeInt},
	}
	// DoctorDoctorToDiagnosisTable holds the schema information for the "doctor_DoctorToDiagnosis" table.
	DoctorDoctorToDiagnosisTable = &schema.Table{
		Name:       "doctor_DoctorToDiagnosis",
		Columns:    DoctorDoctorToDiagnosisColumns,
		PrimaryKey: []*schema.Column{DoctorDoctorToDiagnosisColumns[0], DoctorDoctorToDiagnosisColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "doctor_DoctorToDiagnosis_doctor_id",
				Columns: []*schema.Column{DoctorDoctorToDiagnosisColumns[0]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "doctor_DoctorToDiagnosis_diagnosis_id",
				Columns: []*schema.Column{DoctorDoctorToDiagnosisColumns[1]},

				RefColumns: []*schema.Column{DiagnosesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentResultsTable,
		BloodTypesTable,
		DepartmentsTable,
		DiagnosesTable,
		DoctorsTable,
		GendersTable,
		HospitalsTable,
		MedicalProceduresTable,
		MedicalRecordsTable,
		NursesTable,
		PatientsTable,
		PrefixesTable,
		ProcedureTypesTable,
		RightToTreatmentsTable,
		RightToTreatmentTypesTable,
		RoomsTable,
		TriageResultsTable,
		UrgencyLevelsTable,
		DoctorDoctorToDiagnosisTable,
	}
)

func init() {
	AppointmentResultsTable.ForeignKeys[0].RefTable = DoctorsTable
	AppointmentResultsTable.ForeignKeys[1].RefTable = NursesTable
	AppointmentResultsTable.ForeignKeys[2].RefTable = PatientsTable
	AppointmentResultsTable.ForeignKeys[3].RefTable = RoomsTable
	MedicalProceduresTable.ForeignKeys[0].RefTable = DoctorsTable
	MedicalProceduresTable.ForeignKeys[1].RefTable = PatientsTable
	MedicalProceduresTable.ForeignKeys[2].RefTable = ProcedureTypesTable
	PatientsTable.ForeignKeys[0].RefTable = BloodTypesTable
	PatientsTable.ForeignKeys[1].RefTable = GendersTable
	PatientsTable.ForeignKeys[2].RefTable = PrefixesTable
	RightToTreatmentsTable.ForeignKeys[0].RefTable = HospitalsTable
	RightToTreatmentsTable.ForeignKeys[1].RefTable = PatientsTable
	RightToTreatmentsTable.ForeignKeys[2].RefTable = RightToTreatmentTypesTable
	TriageResultsTable.ForeignKeys[0].RefTable = DepartmentsTable
	TriageResultsTable.ForeignKeys[1].RefTable = NursesTable
	TriageResultsTable.ForeignKeys[2].RefTable = PatientsTable
	TriageResultsTable.ForeignKeys[3].RefTable = UrgencyLevelsTable
	DoctorDoctorToDiagnosisTable.ForeignKeys[0].RefTable = DoctorsTable
	DoctorDoctorToDiagnosisTable.ForeignKeys[1].RefTable = DiagnosesTable
}
