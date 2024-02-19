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

type HabitacionNueva struct {
	Id     int    `json:"Id"`
	Nombre string `json:"Nombre"`
	Camas  int    `json:"Camas"`
	Piezas int    `json:"Piezas"`
}

type Habitaciones struct {
	Habitaciones []HabitacionNueva `json:"Habitaciones"`
}
