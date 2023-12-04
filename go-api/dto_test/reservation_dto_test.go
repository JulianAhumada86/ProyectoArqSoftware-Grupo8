package dto_test

import (
	rDto "go-api/dto/reservations_dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReservationDto (t *testing.T){
	// Crear una instancia del DTO de Booking, si modifico alguna y deja de ser igual, da la alerta
	resDto := rDto.ReservationDto {
		Id: 1,
		InitialDate: "20231221",
		FinalDate: "20231229",
		UserId: 1,
		HotelId: 1,
	}

	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, 1, resDto.Id, "El ID de la reserva no coincide")
	assert.Equal(t, "20231221", resDto.InitialDate, "La fecha de inicio no coincide")
	assert.Equal(t, "20231229", resDto.FinalDate, "La fecha final no coincide")
	assert.Equal(t, 1, resDto.UserId, "El ID de usuario no coincide")
	assert.Equal(t, 1, resDto.HotelId, "El ID de Hotel no coincide")
}