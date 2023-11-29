package image_dto

type ImageDto struct {
	HotelId int `json:"hotel_id"`
	Data    []byte
}
