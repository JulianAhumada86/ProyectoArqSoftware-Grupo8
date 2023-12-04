package services

import (
	hClient "go-api/clients/hotel"
	"log"

	hdto "go-api/dto/hotels_dto"
	e "go-api/errors"

	//hotel_dto "go-api/dto/hotels_dto"

	"go-api/model"
)

type hotelService struct{}

type hotelServicesInterface interface {
	GetHotelbyid(id int) (hdto.HotelDto, error)
	InsertHotel(hotelDto hdto.HotelConHabitaciones) (hdto.HotelConHabitaciones, error)
	GetHotels() (hdto.HotelsDto, e.ErrorApi)
	GetHotelsC() (hdto.HotelsDto, e.ErrorApi)
}

var (
	HotelService hotelServicesInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelbyid(id int) (hdto.HotelDto, error) {

	model_hotel := hClient.GetHotelbyid(id)
	var hotelDto hdto.HotelDto

	hotelDto.Name = model_hotel.Name
	//hotelDto.RoomsAvailable = model_hotel.RoomsAvailable
	hotelDto.Description = model_hotel.Description
	hotelDto.Id = id

	return hotelDto, nil
}

func (s *hotelService) GetHotels() (hdto.HotelsDto, e.ErrorApi) {
	var hotels model.Hotels = hClient.GetHotels()
	hotelsList := make([]hdto.HotelDto, 0)

	for _, hotel := range hotels {
		var hotelDto hdto.HotelDto
		hotelDto.Id = hotel.Id
		hotelDto.Description = hotel.Description
		hotelDto.Name = hotel.Name

		hotelsList = append(hotelsList, hotelDto)
		log.Println(hotel)
	}

	return hdto.HotelsDto{
		Hotels: hotelsList,
	}, nil
}

func (s *hotelService) GetHotelsC() (hdto.HotelsDto, e.ErrorApi) {
	var hotels model.Hotels = hClient.GetHotels()
	hotelsList := make([]hdto.HotelDto, 0)

	for _, hotel := range hotels {
		var hotelDto hdto.HotelDto
		id := hotel.Id
		hotelDto, _ = s.GetHotelbyid(id)
		hotelsList = append(hotelsList, hotelDto)
	}
	return hdto.HotelsDto{
		Hotels: hotelsList,
	}, nil
}

func (s *hotelService) InsertHotel(hotelDto hdto.HotelConHabitaciones) (hdto.HotelConHabitaciones, error) {
	var model_hotel model.Hotel

	model_hotel.Name = hotelDto.Hotel.Name
	model_hotel.Description = hotelDto.Hotel.Description
	model_hotel = hClient.InsertHotel(model_hotel)

	nAmenities := len(hotelDto.Hotel.Amenities)
	newHabitaciones := len(hotelDto.Habitaciones)

	hotelDto.Hotel.Id = model_hotel.Id
	var modelUnionHA model.Hotel_amenitie
	for i := nAmenities; i > 0; i-- {
		modelUnionHA.AmenitieID = hotelDto.Hotel.Amenities[i-1]
		modelUnionHA.HotelID = hotelDto.Hotel.Id
		hClient.AmenitieInHotel(modelUnionHA)

	} //Aca hay que insertar las relaciones
	var modelUnionHH model.Hotel_habitaciones
	for j := newHabitaciones; j > 0; j-- {
		modelUnionHH.Cantidad = hotelDto.Habitaciones[j-1].Cantidad
		modelUnionHH.HabitacionID = hotelDto.Habitaciones[j-1].Id
		modelUnionHH.HotelID = hotelDto.Hotel.Id
		hClient.HabitacionInHotel(modelUnionHH)
		log.Println(hClient.CantHabitaciones(hotelDto.Hotel.Id, hotelDto.Habitaciones[j-1].Id))
	}

	return hotelDto, nil
}
