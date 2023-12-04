package hotel_dto

type Habitacion struct {
	Id       int    `json:"Id"`
	Cantidad int    `json:"Cantidad"`
	Nombre   string `json:"Nombre"`
}

type HotelConHabitaciones struct {
	Hotel        HotelDto     `json:"hotel"`
	Habitaciones []Habitacion `json:"habitaciones"`
}
