package users

import "fmt"

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Status        bool   `json:"status"`
	Creation_date string `json:"creation_date"`
}

var listUsers []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
	LastID() (int, error)
	Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error)
}
type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return listUsers, nil
}
func (r *repository) Store(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	us := User{id, name, last_name, email, age, height, status, creation_date}
	listUsers = append(listUsers, us)
	lastID = us.ID
	return us, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	us := User{Name: name,
		Last_name:     last_name,
		Email:         email,
		Age:           age,
		Height:        height,
		Status:        status,
		Creation_date: creation_date,
	}
	flag := false
	for i := range listUsers {
		if listUsers[i].ID == id {
			us.ID = id
			listUsers[i] = us
			flag = true
		}
	}

	if !flag {
		return User{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return us, nil
}
