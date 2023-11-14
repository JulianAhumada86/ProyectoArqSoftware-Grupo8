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
	InsertHotel(hotelDto hdto.HotelDto) (hdto.HotelDto, error)
	GetHotels() (hdto.HotelsDto, e.ErrorApi)
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
	hotelDto.RoomsAvailable = model_hotel.RoomsAvailable
	hotelDto.Description = model_hotel.Description
	hotelDto.Id = id

	return hotelDto, nil
}

func (s *hotelService) GetHotels() (hdto.HotelsDto, e.ErrorApi) {
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

func (s *hotelService) InsertHotel(hotelDto hdto.HotelDto) (hdto.HotelDto, error) {
	var model_hotel model.Hotel

	model_hotel.Name = hotelDto.Name
	model_hotel.Description = hotelDto.Description
	model_hotel.RoomsAvailable = hotelDto.RoomsAvailable
	log.Println(hotelDto.Description)
	model_hotel = hClient.InsertHotel(model_hotel)

	hotelDto.Id = model_hotel.Id

	return hotelDto, nil
}
