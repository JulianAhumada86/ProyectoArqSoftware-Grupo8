package app

import (
	amenc "go-api/controllers/amenitie"
	hotelc "go-api/controllers/hotel"
	resrc "go-api/controllers/reservation"
	userc "go-api/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	router.POST("/insertHotel/:name/:Nroom/:descr", hotelc.InsertHotel)

	//User
	router.GET("/userId/:id", userc.GetUserById)
	router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email/:Admin", userc.AddUser)
	router.GET("/users", userc.GetUsers)
	router.POST("/login", userc.Login)

	//Reservation
	router.POST("/agregarReservation/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.NewReserva)
	router.GET("/reserva/:id", resrc.GetReservaById)
	router.GET("/reservas", resrc.GetReservas)

	//Amenmitie
	router.POST("/insertAmenitie/:name", amenc.InsertAmenitie)
	router.GET("/getAmenitie/:id", amenc.GetAmenitieById)
	router.GET("/getAmenities", amenc.GetAmenities)

	log.Info("Urls Cargadas")
}
