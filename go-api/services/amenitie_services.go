package services

import (
	hClient "go-api/clients/hotel"
	adto "go-api/dto/amenitie_dto"
	"go-api/model"
)

type amenitiesService struct{}

type amenitiesServicesInterface interface {
	GetAmenitiesbyid(id int) (adto.AmenitieDto, error)
	InsertAmenitie(amenitiDto adto.AmenitieDto) (adto.AmenitieDto, error)
	GetAmenities() (adto.AmenitiesDto, error)
}

var (
	AmenitiesService amenitiesServicesInterface
)

func init() {
	AmenitiesService = &amenitiesService{}
}

func (s *amenitiesService) GetAmenitiesbyid(id int) (adto.AmenitieDto, error) {

	model_amenities := hClient.GetAmenitieById(id)
	var amenitiDto adto.AmenitieDto

	amenitiDto.Id = id
	amenitiDto.Name = model_amenities.Name

	return amenitiDto, nil
}

func (s *amenitiesService) InsertAmenitie(amenitiDto adto.AmenitieDto) (adto.AmenitieDto, error) {
	var model_amenities model.Amenitie
	model_amenities.Name = amenitiDto.Name
	model_amenities = hClient.InsertAmenitie(model_amenities)
	amenitiDto.Id = model_amenities.Id
	return amenitiDto, nil
}

func (s *amenitiesService) GetAmenities() (adto.AmenitiesDto, error) {
	var amenities model.Amenities = hClient.GetAmenities()
	amenitiesList := make([]adto.AmenitieDto, 0)

	for _, amenitie := range amenities {
		var amenitieDto adto.AmenitieDto
		amenitieDto.Id = amenitie.Id
		amenitieDto.Name = amenitie.Name

		amenitiesList = append(amenitiesList, amenitieDto)
	}

	return adto.AmenitiesDto{
		Amenities: amenitiesList,
	}, nil
}
