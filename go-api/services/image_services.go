package services

import (
	iClient "go-api/clients/image"
	image_dto "go-api/dto/image_dto"
	"go-api/model"
)

type imageService struct{}

type imageServicesInterface interface {
	InsertImage(image_dto.ImageDto) (image_dto.ImageDto, error)
	GetImagesByHotelId(id int) (image_dto.ImagesDto, error)
}

var (
	ImageService imageServicesInterface
)

func init() {
	ImageService = &imageService{}
}

func (*imageService) InsertImage(im image_dto.ImageDto) (image_dto.ImageDto, error) {

	_, err := HotelService.GetHotelbyid(im.HotelId)
	if err != nil {
		return im, err
	}
	var mImage model.Image
	mImage.HotelID = im.HotelId
	mImage.Imagen = im.Data

	mImage, err = iClient.InsertImage(mImage)
	if err != nil {
		return im, err
	}

	return im, nil
}

func (*imageService) GetImagesByHotelId(id int) (image_dto.ImagesDto, error) {
	images, _ := iClient.GetImagesByHotelId(id)
	imagesList := make([]image_dto.ImageDto, 0)
	for _, im := range images {
		var dto image_dto.ImageDto
		dto.Data = im.Imagen
		dto.HotelId = im.HotelID
		imagesList = append(imagesList, dto)
	}
	return image_dto.ImagesDto{
		Images: imagesList,
	}, nil
}
