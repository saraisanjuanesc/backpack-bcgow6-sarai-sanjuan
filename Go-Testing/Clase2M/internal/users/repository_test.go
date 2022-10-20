package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	mockeddata []User
}

func (s StubStore) Read(data interface{}) error {
	dat, ok := data.(*[]User)
	if !ok {
		return errors.New("Error")
	}
	*dat = s.mockeddata
	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	esperado := []User{
		{ID: 1,
			Name:          "Juan",
			Last_name:     "Martinez",
			Email:         "juan.martinez@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00"},
		{ID: 2,
			Name:          "Jorge",
			Last_name:     "Salaz",
			Email:         "jorge.salaz@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00"},
	}
	myStubRepository := StubStore{mockeddata: esperado}
	repository := NewRepository(myStubRepository)
	listUs, err := repository.GetAll()

	assert.Equal(t, esperado, listUs)
	assert.Nil(t, err)
}

type MockStore struct {
	ReadCalled bool
	mockeddata []User
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadCalled = true
	dat, ok := data.(*[]User)
	if !ok {
		return errors.New("Error")
	}
	*dat = m.mockeddata

	return nil
}

func (m *MockStore) Write(data interface{}) error {
	return nil
}
func TestUpdateName(t *testing.T) {
	expect_name := "After Update Name "
	expect_lastName := "After Update Last name"

	data := []User{
		{
			ID:            1,
			Name:          "Before Update Name",
			Last_name:     "Before Update Last Name",
			Email:         "juan.martinez@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00"},
	}
	myMockStoreRepository := MockStore{ReadCalled: false, mockeddata: data}

	repository := NewRepository(&myMockStoreRepository)

	resulUser, err := repository.UpdateNameLastName(1, expect_name, expect_lastName)

	assert.Nil(t, err)
	assert.Equal(t, expect_name, resulUser.Name, "No se actualizó el nombre")
	assert.Equal(t, expect_lastName, resulUser.Last_name, "No se actualizó el apellido")
	assert.True(t, myMockStoreRepository.ReadCalled, "No llama a la Función Read")
}

func TestUpdateNameFail(t *testing.T) {
	expect_name := "After Update Name "
	expect_lastName := "After Update Last name"
	myMockStoreRepository := MockStore{ReadCalled: false}

	repository := NewRepository(&myMockStoreRepository)

	resulUser, err := repository.UpdateNameLastName(1, expect_name, expect_lastName)

	assert.Nil(t, err)
	assert.Equal(t, expect_name, resulUser.Name, "No se actualizó el nombre")
	assert.Equal(t, expect_lastName, resulUser.Last_name, "No se actualizó el apellido")
	assert.True(t, myMockStoreRepository.ReadCalled, "No llama a la Función Read")
}
