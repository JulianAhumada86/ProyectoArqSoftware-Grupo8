package services

import (
	hClient "go-api/clients/hotel"
	hdto "go-api/dto/hotels_dto"
	"go-api/model"
)

type amenitiesService struct{}

type amenitiesServicesInterface interface {
	//GetAmenitiesbyid(id int) (hdto.AmenitieDto, error)
	//InsertAmenitie(amenitiDto hdto.AmenitieDto) (hdto.AmenitieDto, error)
	GetAmenities() (hdto.Amenities, error)
}

var (
	AmenitiesService amenitiesServicesInterface
)

func init() {
	AmenitiesService = &amenitiesService{}
}

/*
func (s *amenitiesService) GetAmenitiesbyid(id int) (hdto.AmenitieDto, error) {

	model_amenities := hClient.GetAmenitieById(id)
	var amenitieDto hdto.AmenitieDto

	/*agrego contemplar error*/
/*
	if amenitieDto.Id == 0 {
		return amenitieDto, e.NewBadRequestErrorApi("No se ha encontrado un amenitie")
	}

	amenitieDto.Id = id
	amenitieDto.Name = model_amenities.Name

	return amenitieDto, nil
}

func (s *amenitiesService) InsertAmenitie(amenitiDto hdto.AmenitieDto) (hdto.AmenitieDto, error) {
	var model_amenities model.Amenitie
	model_amenities.Name = amenitiDto.Name
	model_amenities = hClient.InsertAmenitie(model_amenities)
	amenitiDto.Id = model_amenities.Id
	return amenitiDto, nil
}
*/

func (s *amenitiesService) GetAmenities() (hdto.Amenities, error) {
	var amenities model.Amenities = hClient.GetAmenities()
	amenitiesList := make([]hdto.Amenitie, 0)

	for _, amenitie := range amenities {
		var amenitieDto hdto.Amenitie
		amenitieDto.Id = amenitie.Id
		amenitieDto.Name = amenitie.Name

		amenitiesList = append(amenitiesList, amenitieDto)
	}

	return hdto.Amenities{
		Amenities: amenitiesList,
	}, nil
}
