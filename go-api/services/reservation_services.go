package services

import (
	"fmt"
	cl "go-api/clients/reservation"
	reservationDTO "go-api/dto/reservations_dto"
	"go-api/model"
	"time"

	log "github.com/sirupsen/logrus"
)

type reservationService struct{}

type reservationServicesInterface interface {
	NewReserva(reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error)
}

var (
	ReservationService reservationServicesInterface
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) NewReserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error) {
	var Mreserva model.Reservation
	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.Hotel.Id = reserva.HotelId
	log.Debug("reservaAntes:", Mreserva)
	Mreserva.InitialDate = convertirFecha(reserva.InitialDate)

	Mreserva.FinalDate = convertirFecha(reserva.FinalDate)
	Mreserva.User.Id = reserva.UserId

	Mreserva = cl.NewReserva(Mreserva)
	log.Debug("reservaDespues:", Mreserva)
	var rf reservationDTO.ReservationDto

	rf.FinalDate = Mreserva.FinalDate.String()
	rf.HotelName = Mreserva.Hotel.Name
	rf.Id = Mreserva.Id
	rf.InitialDate = Mreserva.InitialDate.String()
	rf.UserName = Mreserva.User.Name

	return rf, nil
}

func convertirFecha(date string) time.Time {
	layout := "2006-01-02"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error al convertir la cadena a tiempo:", err)
		return time.Now()
	} else {
		fmt.Println(t)

	}
	return t
}
