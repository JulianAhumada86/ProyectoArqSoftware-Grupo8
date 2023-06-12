package reservation

import (

	//reservationDTO "go-api/dto/reservations_dto"
	hclient "go-api/clients/hotel"
	uclient "go-api/clients/user"
	"go-api/model"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func NewReserva(reserva model.Reservation) model.Reservation {
	reserva.Hotel = hclient.GetHotelbyid(reserva.Hotel.Id)
	reserva.User = uclient.GetUserById(reserva.User.Id)

	Db.Create(&reserva)

	return reserva
}
