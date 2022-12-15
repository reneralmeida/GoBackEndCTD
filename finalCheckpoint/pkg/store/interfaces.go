package store

import "finalCheckpoint/internal/domain"

type StoreInterface interface {
	ReadDentist(id int) (domain.Dentists, error)
	CreateDentist(dentist domain.Dentists) error
	UpdateDentist(dentist domain.Dentists) error
	DeleteDentist(id int) error
	ExistsDentist(email string) bool
	ReadPatient(id int) (domain.Patients, error)
	CreatePatient(patient domain.Patients) error
	UpdatePatient(patient domain.Patients) error
	DeletePatient(id int) error
	ExistsPatient(document string) bool
	ReadAppointment(id int) (domain.Appointments, error)
	CreateAppointment(appointment domain.Appointments) error
	UpdateAppointment(appointment domain.Appointments) error
	DeleteAppointment(id int) error
	ExistsAppointment(description string) bool
}
