package appointments

import (
	"finalCheckpoint/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Appointments, error)
	Create(p domain.Appointments) (domain.Appointments, error)
	Delete(id int) error
	Update(id int, p domain.Appointments) (domain.Appointments, error)
}

type service struct {
	r Repository
}

func NewAppointmentsService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Appointments, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointments{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Appointments) (domain.Appointments, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Appointments{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Appointments) (domain.Appointments, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointments{}, err
	}
	if u.Description != "" {
		p.Description = u.Description
	}
	if u.Date_And_Time != "" {
		p.Date_And_Time = u.Date_And_Time
	}
	if u.Dentists_Registration > 0 {
		p.Dentists_Registration = u.Dentists_Registration
	}
	if u.Patients_Id > 0 {
		p.Patients_Id = u.Patients_Id
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Appointments{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
