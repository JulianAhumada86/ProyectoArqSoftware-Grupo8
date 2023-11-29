package services

import image_dto "go-api/dto/image_dto"

type imageService struct{}

type imageServicesInterface interface {
	InsertImage(image_dto.ImageDto) (image_dto.ImageDto, error)
}

var (
	ImageService imageServicesInterface
)

func init() {
	ImageService = &imageService{}
}

func (*imageService) InsertImage(im image_dto.ImageDto) (image_dto.ImageDto, error) {

	HotelService.GetHotelbyid(im.HotelId)

	return im, nil
}
