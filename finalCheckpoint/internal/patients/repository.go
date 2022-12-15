package patients

import (
	"errors"
	"fmt"

	"finalCheckpoint/internal/domain"
	"finalCheckpoint/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Patients, error)
	Create(p domain.Patients) (domain.Patients, error)
	Update(id int, p domain.Patients) (domain.Patients, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewPatientsRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r2 *repository) GetByID(id int) (domain.Patients, error) {
	product, err := r2.storage.ReadPatient(id)
	if err != nil {
		return domain.Patients{}, errors.New("Patient not found")
	}
	return product, nil

}

func (r2 *repository) Create(p2 domain.Patients) (domain.Patients, error) {
	if !r2.storage.ExistsPatient(p2.Document) {
		return domain.Patients{}, errors.New("Patient already in the db")
	}
	err := r2.storage.CreatePatient(p2)
	if err != nil {
		return domain.Patients{}, errors.New(fmt.Sprintf("error creating Patient: %s", err.Error()))
	}
	return p2, nil
}

func (r2 *repository) Delete(id int) error {
	err := r2.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}

func (r2 *repository) Update(id int, p2 domain.Patients) (domain.Patients, error) {
	if !r2.storage.ExistsPatient(p2.Document) {
		return domain.Patients{}, errors.New("Patient already in the db")
	}
	err := r2.storage.UpdatePatient(p2)
	if err != nil {
		return domain.Patients{}, errors.New("error updating Patient")
	}
	return p2, nil
}
