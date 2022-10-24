package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll_Repository(t *testing.T) {
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
	myMockStoreRepository := MockStorage{dataMock: esperado}
	repository := NewRepository(&myMockStoreRepository)
	listUs, err := repository.GetAll()

	assert.Equal(t, esperado, listUs)
	assert.Nil(t, err)
}
func TestStore_Repository(t *testing.T) {
	user1 := User{
		ID:            1,
		Name:          "Juan",
		Last_name:     "Martinez",
		Email:         "juan.martinez@mercadolibre.com.mx",
		Age:           27,
		Height:        178,
		Status:        true,
		Creation_date: "2022-05-15T10:12:11+06:00",
	}

	myMockStoreRepository := MockStorage{readCalled: false}
	repository := NewRepository(&myMockStoreRepository)
	resultUser, err := repository.Store(user1.ID, user1.Name, user1.Last_name, user1.Email, user1.Age, user1.Height, user1.Status, user1.Creation_date)

	assert.Nil(t, err)
	assert.True(t, myMockStoreRepository.readCalled, "No llama a la Función Read")
	assert.Equal(t, user1, resultUser, "No retorna el mismo User")
}

func TestLastID_Repository(t *testing.T) {
	expected := 1
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

	myMockStoreRepository := MockStorage{readCalled: false, dataMock: data}
	repository := NewRepository(&myMockStoreRepository)
	lastid, err := repository.LastID()

	assert.Nil(t, err)
	assert.True(t, myMockStoreRepository.readCalled, "No llama a la Función Read")
	assert.Equal(t, expected, lastid, "No retorna el ultimo ID")
}

func TestUpdate_Repository(t *testing.T) {
	userExpected := User{
		ID:            1,
		Name:          "After Update Name",
		Last_name:     "After Update Last Name",
		Email:         "jafter.update@mercadolibre.com.mx",
		Age:           29,
		Height:        167,
		Status:        true,
		Creation_date: "After Update",
	}
	expect_err := User{}
	data := []User{
		{
			ID:            1,
			Name:          "Before Update Name",
			Last_name:     "Before Update Last Name",
			Email:         "before.updatelibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "Before Update"},
	}

	myMockStoreRepository := MockStorage{readCalled: false, dataMock: data}
	repository := NewRepository(&myMockStoreRepository)
	resultUser, err := repository.Update(userExpected.ID, userExpected.Name, userExpected.Last_name, userExpected.Email, userExpected.Age, userExpected.Height, userExpected.Status, userExpected.Creation_date)

	assert.Nil(t, err)
	assert.Equal(t, userExpected, resultUser, "No se actualizó el usuario")

	//ID not found
	resultUser, err = repository.Update(3, userExpected.Name, userExpected.Last_name, userExpected.Email, userExpected.Age, userExpected.Height, userExpected.Status, userExpected.Creation_date)

	assert.NotNil(t, err, "El error no fué encontrado")
	assert.Equal(t, expect_err, resultUser, "No retorna un usuario vacío")
}

func TestUpdateNameLastname_Repository(t *testing.T) {
	expect_name := "After Update Name "
	expect_lastName := "After Update Last name"
	expect_err := User{}

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
	myMockStoreRepository := MockStorage{readCalled: false, dataMock: data}

	repository := NewRepository(&myMockStoreRepository)

	resulUser, err := repository.UpdateNameLastName(1, expect_name, expect_lastName)

	assert.Nil(t, err)
	assert.Equal(t, expect_name, resulUser.Name, "No se actualizó el nombre")
	assert.Equal(t, expect_lastName, resulUser.Last_name, "No se actualizó el apellido")
	assert.True(t, myMockStoreRepository.readCalled, "No llama a la Función Read")

	resulUser, err = repository.UpdateNameLastName(3, expect_name, expect_lastName)
	assert.NotNil(t, err, "El error no fué encontrado")
	assert.Equal(t, expect_err, resulUser, "No retorna un usuario vacío")
}

func TestDelete_Repository(t *testing.T) {
	data := []User{
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

	myMockStoreRepository := MockStorage{dataMock: data, readCalled: false, writeCalled: false}
	repository := NewRepository(&myMockStoreRepository)

	err := repository.Delete(1)

	assert.Nil(t, err, "No se eliminó el usuario")
	assert.True(t, myMockStoreRepository.readCalled, "No llama a la Función Read")
	assert.True(t, myMockStoreRepository.writeCalled, "No llama a la Función Write")

	err = repository.Delete(4)
	assert.NotNil(t, err, "No se encontró el error")
}
