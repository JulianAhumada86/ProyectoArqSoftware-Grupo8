package services

import (
	hClient "go-api/clients/hotel"
	"log"

	hdto "go-api/dto/hotels_dto"
	hotel_dto "go-api/dto/hotels_dto"
	e "go-api/errors"
	"go-api/model"
)

type hotelService struct{}

type hotelServicesInterface interface {
	GetHotelbyid(id int) (hdto.HotelDto, error)
	InsertHotel(hotelDto hdto.HotelDto) (hdto.HotelDto, error)
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

func (s *hotelService) GetHotels(hotel_dto.HotelsDto, e.ErrorApi) {
	var hotels model.Hotel = hClient.GetHotels()
	var hotelsDto hotel_dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto hotel_dto.HotelDto
		id := hotel.Id
		hotelDto, _ = s.GetHotelbyid(id)
		hotelsDto = append(hotelsDto, hotelDto)
	}
	return hotelsDto, nil
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
