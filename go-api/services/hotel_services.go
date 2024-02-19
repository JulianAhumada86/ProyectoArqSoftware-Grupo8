package services

import (
	hClient "go-api/clients/hotel"

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
	AgregarHabitacion(hdto.HabitacionNueva) (hdto.HabitacionNueva, error)
	GetHabitaciones() (hdto.Habitaciones, error)
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
	hotelDto.Description = model_hotel.Description
	hotelDto.Id = id

	for _, habitacionModel := range model_hotel.Habitaciones {
		var habitacionDto hdto.Habitacion
		habitacionDto.Nombre = habitacionModel.Name
		habitacionDto.Id = habitacionModel.Id
		habitacionDto.Cantidad = hClient.CantHabitaciones(model_hotel.Id, habitacionModel.Id)

		hotelDto.Habitaciones = append(hotelDto.Habitaciones, habitacionDto)
	}

	for _, amenitieModel := range model_hotel.Amenities {
		var amenitieDto hdto.Amenitie
		amenitieDto.Id = amenitieModel.Id
		amenitieDto.Name = amenitieModel.Name

		hotelDto.Amenities = append(hotelDto.Amenities, amenitieDto)
	}

	for _, amenitiesModle := range model_hotel.Amenities {
		var amenitieDto hdto.Amenitie
		amenitieDto.Id = amenitiesModle.Id
		amenitieDto.Name = amenitiesModle.Name

		hotelDto.Amenities = append(hotelDto.Amenities, amenitieDto)
	}

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

		for _, habitacion := range hotel.Habitaciones {
			var habitacionDto hdto.Habitacion
			habitacionDto.Nombre = habitacion.Name
			habitacionDto.Id = habitacion.Id
			habitacionDto.Cantidad = hClient.CantHabitaciones(hotel.Id, habitacion.Id)

			hotelDto.Habitaciones = append(hotelDto.Habitaciones, habitacionDto)
		}

		hotelsList = append(hotelsList, hotelDto)
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
		hotelDto.Id = hotel.Id
		hotelDto.Description = hotel.Description
		hotelDto.Name = hotel.Name

		for _, habitacion := range hotel.Habitaciones {
			var habitacionDto hdto.Habitacion
			habitacionDto.Nombre = habitacion.Name
			habitacionDto.Id = habitacion.Id
			habitacionDto.Cantidad = hClient.CantHabitaciones(hotel.Id, habitacion.Id)

			hotelDto.Habitaciones = append(hotelDto.Habitaciones, habitacionDto)
		}

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
		modelUnionHA.AmenitieID = hotelDto.Hotel.Amenities[i-1].Id
		modelUnionHA.HotelID = hotelDto.Hotel.Id
		hClient.AmenitieInHotel(modelUnionHA)

	} //Aca hay que insertar las relaciones
	var modelUnionHH model.Hotel_habitaciones
	for j := newHabitaciones; j > 0; j-- {
		modelUnionHH.Cantidad = hotelDto.Habitaciones[j-1].Cantidad
		modelUnionHH.HabitacionID = hotelDto.Habitaciones[j-1].Id
		modelUnionHH.HotelID = hotelDto.Hotel.Id
		hClient.HabitacionInHotel(modelUnionHH)
		//log.Println(hClient.CantHabitaciones(hotelDto.Hotel.Id, hotelDto.Habitaciones[j-1].Id))
	}

	return hotelDto, nil
}

func (s *hotelService) AgregarHabitacion(habitacion hdto.HabitacionNueva) (hdto.HabitacionNueva, error) {
	var mHabitacion model.Habitacion
	mHabitacion.Name = habitacion.Nombre
	mHabitacion.Camas = habitacion.Camas
	mHabitacion.Piezas = habitacion.Piezas

	mHabitacion = hClient.InsertHabitacion(mHabitacion)

	habitacion.Id = mHabitacion.Id
	return habitacion, nil
}

func (s *hotelService) GetHabitaciones() (hdto.Habitaciones, error) {
	modelHabitaciones, err := hClient.GetHabitaciones()
	if err != nil {
		return hdto.Habitaciones{}, err
	}

	habitacionesList := make([]hdto.HabitacionNueva, 0)

	for _, habitacionM := range modelHabitaciones {
		var habitacion hdto.HabitacionNueva
		habitacion.Id = habitacionM.Id
		habitacion.Camas = habitacionM.Camas
		habitacion.Nombre = habitacionM.Name
		habitacion.Piezas = habitacionM.Piezas

		habitacionesList = append(habitacionesList, habitacion)
	}
	return hdto.Habitaciones{
		Habitaciones: habitacionesList,
	}, nil
}
