package hotel_dto

type Habitacion struct {
	Id       int `json:"Id"`
	Cantidad int `json:"Cantidad"`
}

type HotelConHabitaciones struct {
	Hotel        HotelDto     `json:"hotel"`
	Habitaciones []Habitacion `json:"habitaciones"`
}
