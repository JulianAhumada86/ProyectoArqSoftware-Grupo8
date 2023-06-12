package app

import (
	hotelc "go-api/controllers/hotel"
	resrc "go-api/controllers/reservation"
	userc "go-api/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//router.GET("/prueba")
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	router.POST("/insert/:name/:Nroom/:descr", hotelc.InsertHotel)
	router.GET("/userId/:id", userc.GetUserById)
	router.POST("/agReservation/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.NewReserva)
	router.GET("/reserva/:id")
	log.Print("urls Cargadas")
}
