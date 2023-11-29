package app

import (
	amenc "go-api/controllers/amenitie"
	hotelc "go-api/controllers/hotel"
	imagec "go-api/controllers/image"
	resrc "go-api/controllers/reservation"
	userc "go-api/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	routerUsuario := router.Group("/usuario")
	routerUsuario.Use(TokenMiddleware())

	routerAdmin := router.Group("/admin")
	routerAdmin.Use(AdminTokenMiddleware())

	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	routerAdmin.POST("/insertHotel/:name/:Nroom/:descr", hotelc.InsertHotel)
	router.POST("/insertHotel/:name/:Nroom/:descr", hotelc.InsertHotel)

	//Image
	router.POST("/image/:idHotel", imagec.InsertImage)
	router.GET("/getImagesByHotelId/:idHotel", imagec.GetImagesByHotelId)

	//User
	routerAdmin.GET("/userId/:id", userc.GetUserById)
	router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email", userc.AddUser)
	routerAdmin.GET("/users", userc.GetUsers)
	router.POST("/login", userc.Login)

	//Reservation
	routerAdmin.POST("/agregarReservation/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.NewReserva)
	routerUsuario.GET("/reserva/:id", resrc.GetReservaById)
	routerAdmin.GET("/reservas", resrc.GetReservas)
	router.GET("/dispoibilidadDeReserva/:idHotel/:inicio/:final/:idUser/:habitacion", resrc.Dispoibilidad_de_reserva)
	routerUsuario.GET("/reservaByUserId/:user_id", resrc.GetReservasByUserId)

	//Amenmitie
	router.POST("/insertAmenitie/:name", amenc.InsertAmenitie)
	router.GET("/getAmenitie/:id", amenc.GetAmenitieById)
	router.GET("/getAmenities", amenc.GetAmenities)

	log.Info("Urls Cargadas")
}
