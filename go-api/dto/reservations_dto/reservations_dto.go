package reservations_dto

type ReservationDto struct {
	Id          int    `json:"reservation_id"`
	HotelName   string `json:"hotel_name"`
	HotelId     int    `json:"booked_hotel_id"`
	InitialDate string `json:"initial_date"`
	FinalDate   string `json:"final_date"`
	UserName    string `json:"user_name"`
	UserId      int    `json:"user_booked_id"`
	Address     string `json:"hotel_address"`
}

type ReservationsDto struct {
	Reservations []ReservationDto `json:"reservations"`
}
