package reservations_dto

type ReservationDto struct {
	Id          int    `json:"reservation_id"`
	HotelName   string `json:"hotel_name"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
	UserName    string `json:"user_name"`
}

type ReservationsDto struct {
	Reservations []ReservationDto `json:"reservations"`
}
