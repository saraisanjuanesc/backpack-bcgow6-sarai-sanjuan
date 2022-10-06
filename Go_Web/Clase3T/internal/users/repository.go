package users

import (
	"fmt"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Web/Clase3T/pkg/store"
)

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
	Delete(id int) error
	UpdateNameLastName(id int, name string, last_name string) (User, error)
}
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (users []User, err error) {
	err = r.db.Read(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}
func (r *repository) Store(id int, name string, last_name string, email string, age int, height int, status bool, creation_date string) (User, error) {
	var listUsers []User
	err := r.db.Read(&listUsers)
	if err != nil {
		return User{}, err
	}
	us := User{id, name, last_name, email, age, height, status, creation_date}
	listUsers = append(listUsers, us)

	if err = r.db.Write(listUsers); err != nil {
		return User{}, err
	}
	lastID = us.ID
	return us, nil
}

func (r *repository) LastID() (int, error) {
	var listUsers []User
	err := r.db.Read(&listUsers)
	if err != nil {
		return 0, err
	}

	if len(listUsers) == 0 {
		return 0, nil
	}
	fmt.Println(listUsers[len(listUsers)-1].ID)
	return listUsers[len(listUsers)-1].ID, nil
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
	var listUsers []User
	err := r.db.Read(&listUsers)
	if err != nil {
		return User{}, err
	}
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
	if err = r.db.Write(listUsers); err != nil {
		return User{}, err
	}
	return us, nil
}

func (r *repository) Delete(id int) error {
	delete := false
	var listUsers []User
	err := r.db.Read(&listUsers)
	if err != nil {
		return err
	}
	var index int
	for i := range listUsers {
		if listUsers[i].ID == id {
			index = 1
			delete = true
		}
	}
	if !delete {
		return fmt.Errorf("ID: %d no encontrado", id)
	}
	listUsers = append(listUsers[:index], listUsers[index+1:]...)
	if err = r.db.Write(listUsers); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateNameLastName(id int, name string, last_name string) (User, error) {
	var us User
	update := false
	var listUsers []User
	err := r.db.Read(&listUsers)
	if err != nil {
		return User{}, err
	}
	for i := range listUsers {
		if listUsers[i].ID == id {
			listUsers[i].Name = name
			listUsers[i].Last_name = last_name
			us = listUsers[i]
			update = true
		}
	}
	if !update {
		return User{}, fmt.Errorf("ID: %d no encontrado", id)
	}
	if err = r.db.Write(listUsers); err != nil {
		return User{}, err
	}
	return us, nil
}
