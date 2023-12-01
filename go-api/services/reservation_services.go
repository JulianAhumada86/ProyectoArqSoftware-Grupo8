package services

import (
	"fmt"
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
	GetReservasByUserId(int) (reservations_dto.ReservationsDto, e.ErrorApi)
	Disponibilidad_de_reserva(reservationDTO.ReservationCreateDto) error
}

var (
	ReservationService reservationServicesInterface
	Layoutd            = "2006-01-02"
)

func init() {
	ReservationService = &reservationService{}
}

func (s *reservationService) NewReserva(reserva reservationDTO.ReservationCreateDto) (reservationDTO.ReservationDto, error) {
	var rf reservationDTO.ReservationDto
	err := s.Disponibilidad_de_reserva(reserva) //Lo primero que hago es comprobar la disponibildad de la reserva

	if err != nil {
		log.Println(err.Error())
		return rf, err
	}
	var Mreserva model.Reservation
	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.HotelID = reserva.HotelId

	parseInitial, _ := time.Parse(Layoutd, reserva.InitialDate)

	Mreserva.InitialDate = parseInitial
	parseFinal, _ := time.Parse(Layoutd, reserva.FinalDate)

	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

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
	re.Habitacion = c.Habitacion
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

func (s *reservationService) Disponibilidad_de_reserva(reserva reservationDTO.ReservationCreateDto) error {

	var Mreserva model.Reservation

	Mreserva.Habitacion = reserva.Habitacion
	Mreserva.HotelID = reserva.HotelId

	parseInitial, _ := time.Parse(Layoutd, reserva.InitialDate)

	Mreserva.InitialDate = parseInitial
	parseFinal, _ := time.Parse(Layoutd, reserva.FinalDate)

	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

	if parseFinal.Before(parseInitial) {
		return e.NewBadRequestErrorApi("Fecha inicial despues de final")
	}

	var listaDias []time.Time
	for i := Mreserva.InitialDate; i.Before(Mreserva.FinalDate) || i.Equal(Mreserva.FinalDate); i = i.AddDate(0, 0, 1) {
		listaDias = append(listaDias, i)
	}
	hotel, _ := HotelService.GetHotelbyid(Mreserva.HotelID)

	listaReservas, _ := cl.ComprobarReserva(Mreserva)
	conteoDias := make([]int, len(listaDias))

	for c, dia := range listaDias {
		for _, reserva := range listaReservas {
			if reserva.InitialDate.Before(dia.AddDate(0, 0, -1)) && reserva.FinalDate.After(dia.AddDate(0, 0, 1)) {
				conteoDias[c]++
				if conteoDias[c] >= hotel.RoomsAvailable {
					return e.NewConflictErrorApi(fmt.Sprintf("El dia en la posicion %d no hay disponibilidad", c))
				}
			}
		}
	}
	return nil
}

func (s *reservationService) GetReservasByUserId(id int) (reservations_dto.ReservationsDto, e.ErrorApi) {
	var reservas model.Reservations = cl.GetReservasByUserId(id)

	reservasList := make([]reservationDTO.ReservationDto, 0)
	for _, reserva := range reservas {
		var reservaDto reservations_dto.ReservationDto
		id := reserva.Id
		reservaDto = s.GetReservaById(id)
		//log.Println(reservaDto.Habitacion)
		reservasList = append(reservasList, reservaDto)
	}
	return reservations_dto.ReservationsDto{
		Reservations: reservasList,
	}, nil
}
