package services

import (
	cl "go-api/clients/reservation"
	reservationDTO "go-api/dto/reservations_dto"
	"go-api/model"
	"time"
)

type reservationService struct{}

type reservationServicesInterface interface {
	NewReserva(reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error)
	GetReservaById(int) reservationDTO.ReservationDto
}

var (
	ReservationService reservationServicesInterface
	Layoutd            = "02/01/2006"
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) NewReserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error) {
	var Mreserva model.Reservation

	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.HotelID = reserva.HotelId

	parseInitial, _ := time.Parse(Layoutd, reserva.InitialDate)
	Mreserva.InitialDate = parseInitial
	parseFinal, _ := time.Parse(Layoutd, reserva.FinalDate)

	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

	Mreserva = cl.NewReserva(Mreserva)

	var rf reservationDTO.ReservationDto

	rf.FinalDate = Mreserva.FinalDate.String()
	rf.HotelName = Mreserva.Hotel.Name
	rf.Id = Mreserva.Id
	rf.InitialDate = Mreserva.InitialDate.String()
	rf.UserName = Mreserva.User.Name

	return rf, nil
}

func (s *reservationService) GetReservaById(id int) reservationDTO.ReservationDto {
	var re reservationDTO.ReservationDto

	re.Id = id

	c := cl.GetReservaById(id)

	re.FinalDate = c.FinalDate.Format(Layoutd)
	re.HotelName = c.Hotel.Name
	re.Id = c.Id
	re.InitialDate = c.InitialDate.Format(Layoutd)
	re.UserName = (c.User.Name + " " + c.User.LastName)

	return re
}
