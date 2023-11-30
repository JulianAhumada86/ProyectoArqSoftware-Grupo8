package tests

import (
	rdto "go-api/dto/reservations_dto"
	"go-api/errors"
	"net/http"
)

type TestReservations struct {}

func (t *TestReservations) InsertBooking (resDto rdto.ReservationDto)(rdto.ReservationDto, errors.ErrorApi){
	if resDto.UserId == 0 {
		return rdto.ReservationDto{}, errors.NewErrorApi("Error al insertar la reserva", "reservation_insert_error", http.StatusInternalServerError, nil)
	}

	return rdto.ReservationDto{}, nil
}

func (t *TestReservations) GetReservationByUserId(id int) (rdto.ReservationDto, errors.ErrorApi) {
	if id == 1 {
		return rdto.ReservationDto {
			Id: 1,
			InitialDate: 20230418,
			FinalDate: 20220126,
			UserId: 1,
			HotelId: 2,
		}, nil
	}
	return rdto.ReservationDto{}, errors.NewNotFoundErrorApi("Resera no encontrada")
}