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
func ComprobarReserva(reserva model.Reservation) (model.Reservations, error) {
	var reservas model.Reservations
	hId := reserva.HotelID
	habitacion := reserva.Habitacion
	inicio := reserva.InitialDate
	final := reserva.FinalDate
	final = final.AddDate(0, 0, 1) //Agrego un dia para comprobar el ultimo dia en el bucle de abajo

	Db.Where("? < final_date AND hotel_id = ? AND habitacion = ?", inicio, hId, habitacion).Find(&reservas)

	return reservas, nil
}

func GetReservas() model.Reservations {
	var reservas model.Reservations
	Db.Find(&reservas)

	return reservas
}
func GetReservasByUserId(id int) model.Reservations {
	var reservas model.Reservations
	Db.Where("user_id = ?", id).Preload("Hotel").Preload("User").Find(&reservas)

	return reservas
}

func CantHabitaciones(idH int, idA int) int {
	var x model.Hotel_habitaciones
	Db.Where("hotel_id = ? and habitacion_id = ?", idH, idA).First(&x)

	return x.Cantidad
}
