package app

import (
	amenc "go-api/controllers/amenitie"
	hotelc "go-api/controllers/hotel"
	resrc "go-api/controllers/reservation"
	userc "go-api/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	routerVerificado := router.Group("/usuario")
	routerVerificado.Use(TokenMiddleware())

	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	routerVerificado.POST("/insertHotel/:name/:Nroom/:descr", hotelc.InsertHotel)

	//User
	routerVerificado.GET("/userId/:id", userc.GetUserById)
	router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email/:Admin", userc.AddUser)
	routerVerificado.GET("/users", userc.GetUsers)
	router.POST("/login", userc.Login)

	//Reservation
	routerVerificado.POST("/agregarReservation/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.NewReserva)
	routerVerificado.GET("/reserva/:id", resrc.GetReservaById)
	routerVerificado.GET("/reservas", resrc.GetReservas)
	router.GET("/dispoibilidadDeReserva/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.Dispoibilidad_de_reserva)

	//Amenmitie
	router.POST("/insertAmenitie/:name", amenc.InsertAmenitie)
	router.GET("/getAmenitie/:id", amenc.GetAmenitieById)
	router.GET("/getAmenities", amenc.GetAmenities)

	log.Info("Urls Cargadas")
}
