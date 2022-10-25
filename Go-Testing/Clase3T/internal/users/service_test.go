package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []User{
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

func TestGetAll_Service(t *testing.T) {
	myMockStorage := MockStorage{DataMock: data, readCalled: false}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	resultUser, err := service.GetAll()

	assert.Nil(t, err, "")
	assert.Equal(t, data, resultUser, "")

}

func TestStore_Service(t *testing.T) {
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
	myMockStorage := MockStorage{readCalled: false, writeCalled: false}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	resultUser, err := service.Store(user1.Name, user1.Last_name, user1.Email, user1.Age, user1.Height, user1.Status, user1.Creation_date)

	assert.Nil(t, err)
	assert.Equal(t, user1, resultUser)

}

func TestUpdate_Service(t *testing.T) {
	userExpectederr := User{}
	userExpected := User{
		ID:            1,
		Name:          "Juan Pablo",
		Last_name:     "Martinez Gonzalez",
		Email:         "juan.gonzalez@mercadolibre.com.mx",
		Age:           29,
		Height:        188,
		Status:        true,
		Creation_date: "2022-08-17T10:12:11+06:00",
	}

	myMockStorage := MockStorage{DataMock: data}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	userUpdated, err := service.Update(userExpected.ID, userExpected.Name, userExpected.Last_name, userExpected.Email, userExpected.Age, userExpected.Height, userExpected.Status, userExpected.Creation_date)
	assert.Nil(t, err)
	assert.Equal(t, userExpected, userUpdated, "Los usuarios deben ser iguales")
	assert.True(t, myMockStorage.readCalled, "No llama a la función Read del Storage")

	userUpdated, err = service.Update(3, userExpected.Name, userExpected.Last_name, userExpected.Email, userExpected.Age, userExpected.Height, userExpected.Status, userExpected.Creation_date)
	assert.NotNil(t, err)
	assert.Equal(t, userExpectederr, userUpdated)

}

func TestDelete_Service(t *testing.T) {
	myMockStorage := MockStorage{DataMock: data}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	err := service.Delete(1)

	assert.Nil(t, err, "%s", err)
	assert.True(t, myMockStorage.readCalled, "No llama a la función Read del Storage")
	assert.True(t, myMockStorage.writeCalled, "No llama a la función Write del Storage")

	err = service.Delete(1)
	assert.NotNil(t, err)
}

func TestUpdateNameLastName_Service(t *testing.T) {
	expect_name := "After Update Name "
	expect_lastName := "After Update Last name"
	expect_err := User{}

	datas := []User{
		{ID: 1,
			Name:          "Juan",
			Last_name:     "Martinez",
			Email:         "juan.martinez@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00"}}
	myMockStorage := MockStorage{DataMock: datas, readCalled: false}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)
	resulUser, err := service.UpdateNameLastName(1, expect_name, expect_lastName)

	assert.Nil(t, err)
	assert.Equal(t, expect_name, resulUser.Name, "No se actualizó el nombre")
	assert.Equal(t, expect_lastName, resulUser.Last_name, "No se actualizó el apellido")
	assert.True(t, myMockStorage.readCalled, "No llama a la Función Read")

	resulUser, err = repository.UpdateNameLastName(3, expect_name, expect_lastName)
	assert.NotNil(t, err, "El error no fué encontrado")
	assert.Equal(t, expect_err, resulUser, "No retorna un usuario vacío")

}
