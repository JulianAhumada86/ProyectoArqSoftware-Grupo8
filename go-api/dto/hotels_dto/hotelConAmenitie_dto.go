package hotel_dto

type Amenitie struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Amenities struct {
	Amenities []Amenitie `json:"amenities"`
}
