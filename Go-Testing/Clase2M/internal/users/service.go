package users

import "fmt"

type Service interface {
	GetAll() ([]User, error)
	Store(name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
	Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
	Delete(id int) error
	UpdateNameLastName(id int, name string, last_name string) (User, error)
}

type service struct {
	repository Repository
}

func NewServices(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]User, error) {
	list, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (s *service) Store(name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	lastid, err := s.repository.LastID()
	fmt.Println(lastid)
	if err != nil {
		return User{}, nil
	}
	lastid++

	newUser, err := s.repository.Store(lastid, name, last_name, email, age, height, status, creation_date)
	if err != nil {
		return User{}, nil
	}
	return newUser, nil
}

func (s *service) Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	return s.repository.Update(id, name, last_name, email, age, height, status, creation_date)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNameLastName(id int, name string, last_name string) (User, error) {
	return s.repository.UpdateNameLastName(id, name, last_name)
}
