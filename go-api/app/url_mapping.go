package app

import (
	hotelc "go-api/controllers/hotel"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//router.GET("/prueba")
	router.GET("/prueba/:id", hotelc.GetHotelbyid)
	router.GET("/insert/:name/:Nroom/:descr", hotelc.InsertHotel)
	log.Print("Hola")
}
