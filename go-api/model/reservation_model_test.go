package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservatino(t *testing.T){
	assert := assert.New(t)

	user := User{
		Id: 1,
		Name: "test",
		LastName: "test",
		DNI: "test",
		Email: "test@test.com",
		Password: "test",
		Type: false,
	}

	hotel := Hotel{
		Id: 1,
		Name: "Test Hotel",
		RoomsAvailable: 10,
		Description: "Test Description Hotel",
		//aca deberiamos poner amenitites y fotos creo
	}

	reservation := Reservation{
		Id: 1,
		//InitialDate: "20231221",
		//FinalDate: "20231229",
		Habitacion: "test habitacion",
		UserID: 1,
		HotelID: 1,
		User: user,
		Hotel: hotel,
	}

	expectedDateFrom := "20231221"
	assert.Equal(expectedDateFrom, reservation.InitialDate, "Se espera que la fecha de inicio sea %v", expectedDateFrom)

	expectedDateTo := "20231229"
	assert.Equal(expectedDateTo, reservation.FinalDate, "Se espero que la fecha de fina sea %v", expectedDateTo)

	assert.Equal(1, reservation.Id, "El ID de la reserva no coincide")

	assert.Equal("test", reservation.User.Name, "El nombre no coincide")
	assert.Equal("test", reservation.User.Email, "El email no coincide")

	assert.Equal("Test Hotel", reservation.Hotel.Name, "El nombre del hotel no coincide")
	assert.Equal(10, reservation.Hotel.RoomsAvailable, "La cantidad de habitaciones no coincide")
} 