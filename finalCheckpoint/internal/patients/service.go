package patients

import (
	"finalCheckpoint/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Patients, error)
	Create(p domain.Patients) (domain.Patients, error)
	Delete(id int) error
	Update(id int, p domain.Patients) (domain.Patients, error)
}

type service struct {
	r2 Repository
}

func NewPatientsService(r2 Repository) Service {
	return &service{r2}
}

func (s2 *service) GetByID(id int) (domain.Patients, error) {
	p2, err := s2.r2.GetByID(id)
	if err != nil {
		return domain.Patients{}, err
	}
	return p2, nil
}

func (s2 *service) Create(p2 domain.Patients) (domain.Patients, error) {
	p2, err := s2.r2.Create(p2)
	if err != nil {
		return domain.Patients{}, err
	}
	return p2, nil
}
func (s2 *service) Update(id int, u2 domain.Patients) (domain.Patients, error) {
	p2, err := s2.r2.GetByID(id)
	if err != nil {
		return domain.Patients{}, err
	}
	if u2.Name != "" {
		p2.Name = u2.Name
	}
	if u2.Lastname != "" {
		p2.Lastname = u2.Lastname
	}
	if u2.Document != "" {
		p2.Document = u2.Document
	}
	if u2.Reg_Date != "" {
		p2.Reg_Date = u2.Reg_Date
	}
	p2, err = s2.r2.Update(id, p2)
	if err != nil {
		return domain.Patients{}, err
	}
	return p2, nil
}

func (s2 *service) Delete(id int) error {
	err := s2.r2.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
