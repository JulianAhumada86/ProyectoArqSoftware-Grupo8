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
	Mreserva.HabitacionId = reserva.HabitacionId
	Mreserva.HotelID = reserva.HotelId

	parseInitial, err := time.Parse(Layoutd, reserva.InitialDate)

	if err != nil {
		log.Fatalf(err.Error())
	}
	Mreserva.InitialDate = parseInitial
	parseFinal, err := time.Parse(Layoutd, reserva.FinalDate)
	if err != nil {
		log.Fatalf(err.Error())
	}
	Mreserva.FinalDate = parseFinal
	Mreserva.UserID = reserva.UserId

	Mreserva = cl.NewReserva(Mreserva)

	rf.FinalDate = Mreserva.FinalDate.String()
	rf.Habitacion = Mreserva.Habitacion.Id
	rf.HotelId = Mreserva.Hotel.Id
	rf.HotelName = Mreserva.Hotel.Name
	rf.Id = Mreserva.Id
	rf.InitialDate = Mreserva.InitialDate.String()
	rf.UserName = Mreserva.User.Name
	rf.UserId = Mreserva.User.Id
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
	//c.HabitacionId
	//Poner un buscar habitacion por id
	re.Habitacion = c.HabitacionId
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
	Mreserva.HabitacionId = reserva.HabitacionId
	Mreserva.Habitacion.Id = reserva.HabitacionId
	Mreserva.Hotel.Id = reserva.HotelId
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

	cantHabitaciones := cl.CantHabitaciones(Mreserva.HotelID, Mreserva.Habitacion.Id)
	listaReservas, _ := cl.ComprobarReserva(Mreserva)

	conteoDias := make([]int, len(listaDias))
	for c, dia := range listaDias {
		for _, reserva := range listaReservas {
			if reserva.InitialDate.Before(dia.AddDate(0, 0, -1)) && reserva.FinalDate.After(dia.AddDate(0, 0, 1)) {
				conteoDias[c]++
				if conteoDias[c] >= cantHabitaciones {
					return e.NewBadRequestErrorApi(fmt.Sprintf("El dia %d/%d/%d no hay disponibilidad", listaDias[c].Day(), int(listaDias[c].Month()), listaDias[c].Year()))

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
