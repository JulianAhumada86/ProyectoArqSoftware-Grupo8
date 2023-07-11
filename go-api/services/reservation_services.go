package services

import (
	cl "go-api/clients/reservation"
	"go-api/dto/reservations_dto"
	reservationDTO "go-api/dto/reservations_dto"
	e "go-api/errors"
	"go-api/model"
	"log"
	"time"
)

type reservationService struct{}

type reservationServicesInterface interface {
	NewReserva(reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error)
	GetReservaById(int) reservationDTO.ReservationDto
	GetReservas() (reservations_dto.ReservationsDto, e.ErrorApi)
	//GetReservaByUserId(id int) (reservations_dto.ReservationDto, e.ErrorApi)
	Disponibilidad_de_reserva(reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error)
}

var (
	ReservationService reservationServicesInterface
	Layoutd            = "2006-01-02"
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) NewReserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error) {
	var Mreserva model.Reservation
	var rf reservationDTO.ReservationDto
	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.HotelID = reserva.HotelId

	parseInitial, _ := time.Parse(Layoutd, reserva.InitialDate)

	Mreserva.InitialDate = parseInitial
	parseFinal, _ := time.Parse(Layoutd, reserva.FinalDate)

	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

	if parseFinal.Before(parseInitial) {

		return rf, e.NewBadRequestErrorApi("Fecha inicial antes de la final")
	}

	if cl.ComprobarReserva(Mreserva) {
		Mreserva = cl.NewReserva(Mreserva)
		log.Println("Esta Disponible")

	} else {
		log.Println("No esta disponible")
		return rf, e.NewBadRequestErrorApi("Cosas") //Completar este error bien, no me acuerdo como era
	}

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

func (s *reservationService) GetReservas() (reservations_dto.ReservationsDto, e.ErrorApi) {
	var reservas model.Reservations = cl.GetReservas()

	reservasList := make([]reservationDTO.ReservationDto, 0)
	for _, reserva := range reservas {
		var reservaDto reservations_dto.ReservationDto
		id := reserva.Id
		reservaDto = s.GetReservaById(id)

		reservasList = append(reservasList, reservaDto)
	}
	return reservations_dto.ReservationsDto{
		Reservations: reservasList,
	}, nil
}

func (s *reservationService) Disponibilidad_de_reserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error) {

	var Mreserva model.Reservation
	var rf reservationDTO.ReservationDto
	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.HotelID = reserva.HotelId

	parseInitial, _ := time.Parse(Layoutd, reserva.InitialDate)

	Mreserva.InitialDate = parseInitial
	parseFinal, _ := time.Parse(Layoutd, reserva.FinalDate)

	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

	if parseFinal.Before(parseInitial) {
		return rf, e.NewBadRequestErrorApi("Fecha inicial antes de la final")
	}

	if cl.ComprobarReserva(Mreserva) {
		return rf, nil

	} else {
		return rf, e.NewBadRequestErrorApi("Las fechas no estan disponibles") //Completar este error bien, no me acuerdo como era
	}

}

/*
func (s *reservationService) GetReservaByUserId(id int)(reservationDTO.ReservationDto, e.ErrorApi){
	var reserva model.Reservation = cl.GetReservaByUserId(id)
	var reservaDto reservations_dto.ReservationDto

	if reserva.Id == 0 {
		return reservaDto, e.NewBadRequestErrorApi("Reserva no encontrada")
	}

	reservaDto.Id = reserva.Id
	reservaDto.InitialDate = reserva.InitialDate
	reservaDto.FinalDate = reserva.FinalDate
	reservaDto.HotelId = reserva.HotelID
	reservaDto.HotelName = reserva.Hotel
	reservaDto.Address = reserva.

	return reservaDto, nil
}
*/
