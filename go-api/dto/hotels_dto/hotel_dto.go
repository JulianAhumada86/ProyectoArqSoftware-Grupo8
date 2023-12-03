package hotel_dto

type HotelDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amenities   []int  `json:"amenities"`
}

type HotelsDto struct {
	Hotels []HotelDto `json:"hotels"`
}
