package users

type Service interface {
	GetAll() ([]User, error)
	Store(name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
	Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
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
	if err != nil {
		return User{}, nil
	}
	lastID++

	newUser, err := s.repository.Store(lastid, name, last_name, email, age, height, status, creation_date)
	if err != nil {
		return User{}, nil
	}
	lastID++
	return newUser, nil
}

func (s *service) Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	return s.repository.Update(id, name, last_name, email, age, height, status, creation_date)
}
