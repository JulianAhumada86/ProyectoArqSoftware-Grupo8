package hotel_dto

type HotelDto struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Amenities      []string `json:"amenities"`
	RoomsAvailable int      `json:"rooms_available"`
}

type HotelsDto struct {
	Hotels []HotelDto `json:"hotels"`
}
