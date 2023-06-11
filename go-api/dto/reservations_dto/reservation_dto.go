package reservation_dto

type ReservationDto struct {
	Id          int    `json:"id"`
	HotelName   string `json:"hotel_name"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
}

type ReservationsDto struct {
	Reservations []ReservationDto `json:"reservations"`
}
