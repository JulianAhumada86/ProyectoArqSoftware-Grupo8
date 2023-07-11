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
func ComprobarReserva(reserva model.Reservation) bool {
	var reservas model.Reservations
	hId := reserva.HotelID
	habitacion := reserva.Habitacion
	inicio := reserva.InitialDate
	final := reserva.FinalDate
	final = final.AddDate(0, 0, 1) //Agrego un dia para comprobar el ultimo dia en el bucle de abajo

	for inicio != final {
		Db.Where("? BETWEEN initial_date AND final_date AND hotel_id = ? AND habitacion = ?", inicio, hId, habitacion).Find(&reservas)
		if len(reservas) >= 4 {
			return false //Esto quiere decir que no hay mas dispoibilidad, el signo > es por las dudas un error de la base de datos
		}
		inicio = inicio.AddDate(0, 0, 1)
	}

	return true
}

func GetReservas() model.Reservations {
	var reservas model.Reservations
	Db.Find(&reservas)

	return reservas
}
