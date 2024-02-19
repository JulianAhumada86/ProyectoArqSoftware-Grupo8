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
	//Comprueba que este logeado
	routerUsuario := router.Group("/usuario")
	routerUsuario.Use(TokenMiddleware())

	//Comprueba si esta logueado y es admin
	routerAdmin := router.Group("/admin")
	routerAdmin.Use(AdminTokenMiddleware())

	//Habitacion

	//Hotel
	router.GET("/hotelId/:id", hotelc.GetHotelbyid)
	router.GET("/hotels", hotelc.GetHotelsC)
	routerAdmin.POST("/insertHotel", hotelc.InsertHotel)
	routerAdmin.GET("/hotels", hotelc.GetHotels)

	routerAdmin.POST("/AgregarHabitacion/:name/:camas/:piezas", hotelc.AgregarHabitacion)
	routerAdmin.GET("/Habitaciones", hotelc.GetHabitaciones)

	//Image
	router.GET("/getImagesByHotelId/:idHotel", imagec.GetImagesByHotelId)
	routerAdmin.POST("/image/:idHotel", imagec.InsertImage)

	//User
	router.POST("/addUsuario/:name/:LastName/:DNI/:Password/:Email", userc.AddUser)
	router.POST("/login", userc.Login)
	routerUsuario.GET("/userId/", userc.GetUserById)
	routerAdmin.GET("/userId/:idUsuario", userc.GetUserById)
	routerAdmin.GET("/users", userc.GetUsers)

	//Reservation
	routerUsuario.POST("/agregarReservation/:idHotel/:inicio/:final/:habitacion", resrc.NewReserva)
	routerUsuario.GET("/disponibilidadDeReserva/:idHotel/:inicio/:final/:habitacion", resrc.Disponibilidad_de_reserva)
	routerUsuario.GET("/reserva/:id", resrc.GetReservaById)
	routerUsuario.GET("/reservaByUserId/:user_id", resrc.GetReservasByUserId)
	routerAdmin.GET("/reservas", resrc.GetReservas)

	router.GET("/disponibilidadDeReserva/:idHotel/:inicio/:final/:habitacion", resrc.Disponibilidad_de_reserva)

	//Amenmitie
	router.POST("/insertAmenitie/:name", amenc.InsertAmenitie)
	router.GET("/getAmenitie/:id", amenc.GetAmenitieById)
	router.GET("/getAmenities", amenc.GetAmenities)

	//borrar al terminar
	router.POST("/image/:idHotel", imagec.InsertImage)
	router.POST("/insertHotel", hotelc.InsertHotel)

	log.Info("Urls Cargadas")
}
