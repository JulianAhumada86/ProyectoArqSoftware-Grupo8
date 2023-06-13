package reservation

import (

	//reservationDTO "go-api/dto/reservations_dto"

	"go-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func NewReserva(reserva model.Reservation) model.Reservation {

	Db.Create(&reserva)
	log.Debug("reserva culea:", reserva)
	return reserva
}

func GetReservaById(id int) model.Reservation {
	var reserva model.Reservation

	Db.Where("id = ?", id).Preload("Hotel").Preload("User").First(&reserva)
	log.Debug("reserva:", reserva)
	return reserva

}

func GetReservas() model.Reservations {
	var reservas model.Reservations
	Db.Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}
