package dentists

import (
	"errors"
	"fmt"

	"finalCheckpoint/internal/domain"
	"finalCheckpoint/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Dentists, error)
	Create(p domain.Dentists) (domain.Dentists, error)
	Update(id int, p domain.Dentists) (domain.Dentists, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Dentists, error) {
	product, err := r.storage.ReadDentist(id)
	if err != nil {
		return domain.Dentists{}, errors.New("dentist not found")
	}
	return product, nil

}

func (r *repository) Create(p domain.Dentists) (domain.Dentists, error) {
	if !r.storage.ExistsDentist(p.Email) {
		return domain.Dentists{}, errors.New("dentist already in the db")
	}
	err := r.storage.CreateDentist(p)
	if err != nil {
		return domain.Dentists{}, errors.New(fmt.Sprintf("error creating dentist: %s", err.Error()))
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Dentists) (domain.Dentists, error) {
	if !r.storage.ExistsDentist(p.Email) {
		return domain.Dentists{}, errors.New("code value already exists")
	}
	err := r.storage.UpdateDentist(p)
	if err != nil {
		return domain.Dentists{}, errors.New("error updating dentist")
	}
	return p, nil
}
