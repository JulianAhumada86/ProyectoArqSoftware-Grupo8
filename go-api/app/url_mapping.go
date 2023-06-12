package app

import (
	hotelc "go-api/controllers/hotel"
	userc "go-api/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//router.GET("/prueba")
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	router.POST("/insert/:name/:Nroom/:descr", hotelc.InsertHotel)
	router.GET("/userId/:id", userc.GetUserById)
	log.Print("urls Cargadas")
}
