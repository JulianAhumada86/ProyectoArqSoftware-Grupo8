package service_test

import (
	rDto "go-api/dto/reservations_dto"
	e "go-api/errors"
	"go-api/services"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestReservations struct {}

func (t *TestReservations) NewReservation(resDto rDto.ReservationCreateDto)(rDto.ReservationDto, e.ErrorApi){
	if resDto.UserId == 0 {
		return rDto.ReservationDto{}, e.NewErrorApi("Error al insertar la reserva", "booking_insert_error", http.StatusInternalServerError, nil)
	}

	return rDto.ReservationDto{}, nil
}

func (t *TestReservations) GetReservationById(id int)(rDto.ReservationDto, e.ErrorApi){
	return rDto.ReservationDto{}, nil 
}

func (t *TestReservations) GetReservation()(rDto.ReservationsDto, e.ErrorApi){
	return rDto.ReservationsDto{}, nil 
}

func GetReservationsByUserId(id int)(rDto.ReservationsDto, e.ErrorApi){
	return rDto.ReservationsDto{}, nil
}


/*
func TestNewReservation(t *testing.T){
	res := rDto.ReservationCreateDto{
		UserId: 1,
		HotelId: 1,
		InitialDate: "20231221",
		FinalDate: "20231229",
		Habitacion: "Habitacion 1",
	}

	services.ReservationService = &TestReservations{}

	//ACA QUEDARIA HACER EN BASE A LA DISPONIBILIDAD
}
*/

func TestGetReservationsByUserId(t *testing.T){
		// Configuramos el servicio con un mock
		services.ReservationService = &TestReservations{}

		// Datos esperados para la reserva del usuario
		expectedReservations := rDto.ReservationsDto{
			Reservations: []rDto.ReservationDto{
				{
					Id:         1,
					InitialDate: "20231221",
					FinalDate:   "20231229",
					UserId:      1,
					HotelId:     1,
				},
			},
		}
	
		// Llamamos a la función del servicio que queremos probar
		catchedReservations, err := services.ReservationService.GetReservasByUserId(expectedReservations.Reservations[0].UserId)
	
		// Verificamos si hay algún error
		assert.Nil(t, err)
	
		// Comparamos los resultados obtenidos con los esperados
		assert.Len(t, catchedReservations.Reservations, len(expectedReservations.Reservations), "La cantidad de reservas no coincide")
	
		for i, expected := range expectedReservations.Reservations {
			assert.Equal(t, expected.Id, catchedReservations.Reservations[i].Id, "El ID de reserva no coincide")
			assert.Equal(t, expected.InitialDate, catchedReservations.Reservations[i].InitialDate, "La fecha de inicio no coincide")
			assert.Equal(t, expected.FinalDate, catchedReservations.Reservations[i].FinalDate, "La fecha de fin no coincide")
			assert.Equal(t, expected.UserId, catchedReservations.Reservations[i].UserId, "El ID de usuario no coincide")
			assert.Equal(t, expected.HotelId, catchedReservations.Reservations[i].HotelId, "El ID de hotel no coincide")
		}
}
