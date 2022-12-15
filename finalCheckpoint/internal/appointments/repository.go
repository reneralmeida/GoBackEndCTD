package appointments

import (
	"errors"
	"fmt"

	"finalCheckpoint/internal/domain"
	"finalCheckpoint/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Appointments, error)
	Create(p domain.Appointments) (domain.Appointments, error)
	Update(id int, p domain.Appointments) (domain.Appointments, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewAppointmentsRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Appointments, error) {
	product, err := r.storage.ReadAppointment(id)
	if err != nil {
		return domain.Appointments{}, errors.New("Appointment not found")
	}
	return product, nil

}

func (r *repository) Create(p domain.Appointments) (domain.Appointments, error) {
	if !r.storage.ExistsAppointment(p.Description) {
		return domain.Appointments{}, errors.New("Appointment already in the db")
	}
	err := r.storage.CreateAppointment(p)
	if err != nil {
		return domain.Appointments{}, errors.New(fmt.Sprintf("error creating Appointment: %s", err.Error()))
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Appointments) (domain.Appointments, error) {
	if !r.storage.ExistsAppointment(p.Description) {
		return domain.Appointments{}, errors.New("Appointment already in the db")
	}
	err := r.storage.UpdateAppointment(p)
	if err != nil {
		return domain.Appointments{}, errors.New("error updating Appointment")
	}
	return p, nil
}
