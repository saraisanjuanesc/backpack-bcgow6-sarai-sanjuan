package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {

	user := User{
		ID:            1,
		Name:          "Juan",
		Last_name:     "Martinez",
		Email:         "juan.martinez@mercadolibre.com.mx",
		Age:           27,
		Height:        178,
		Status:        true,
		Creation_date: "2022-05-15T10:12:11+06:00",
	}
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

	myMockStorage := MockStorage{dataMock: []User{user}}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	userUpdated, err := service.Update(userExpected.ID, userExpected.Name, userExpected.Last_name, userExpected.Email, userExpected.Age, userExpected.Height, userExpected.Status, userExpected.Creation_date)
	assert.Nil(t, err)
	assert.Equal(t, userExpected, userUpdated, "Los usuarios deben ser iguales")
	assert.True(t, myMockStorage.readCalled, "No llama a la función Read del Storage")

}

func TestDelete(t *testing.T) {
	users := []User{
		{
			ID:            1,
			Name:          "Juan",
			Last_name:     "Martinez",
			Email:         "juan.martinez@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00",
		},
		{
			ID:            2,
			Name:          "Jorge",
			Last_name:     "Salaz",
			Email:         "jorge.salaz@mercadolibre.com.mx",
			Age:           27,
			Height:        178,
			Status:        true,
			Creation_date: "2022-05-15T10:12:11+06:00",
		},
	}

	myMockStorage := MockStorage{dataMock: users}
	repository := NewRepository(&myMockStorage)
	service := NewServices(repository)

	err := service.Delete(1)

	assert.Nil(t, err, "%s", err)
	assert.True(t, myMockStorage.readCalled, "No llama a la función Read del Storage")
	assert.True(t, myMockStorage.writeCalled, "No llama a la función Write del Storage")

	err = service.Delete(1)
	assert.Nil(t, err, "%s", err)
	assert.True(t, myMockStorage.readCalled, "No llama a la función Read del Storage")
	assert.True(t, myMockStorage.writeCalled, "No llama a la función Write del Storage")
}
