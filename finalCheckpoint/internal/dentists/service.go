package dentists

import (
	"finalCheckpoint/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentists, error)
	Create(p domain.Dentists) (domain.Dentists, error)
	Delete(id int) error
	Update(id int, p domain.Dentists) (domain.Dentists, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Dentists, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentists{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Dentists) (domain.Dentists, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentists{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Dentists) (domain.Dentists, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentists{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Lastname != "" {
		p.Lastname = u.Lastname
	}
	if u.Registration > 0 {
		p.Registration = u.Registration
	}
	if u.Email != "" {
		p.Email = u.Email
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Dentists{}, err
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
