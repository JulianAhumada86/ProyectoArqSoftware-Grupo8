package reservation

import (
	reservationDTO "go-api/dto/reservations_dto"
	se "go-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewReserva(ctx *gin.Context) {
	var create reservationDTO.ReservationCreateDto

	idH, _ := strconv.Atoi(ctx.Param("idHotel"))
	inicio := ctx.Param("inicio")
	final := ctx.Param("final")
	idU, _ := strconv.Atoi(ctx.Param("idUser"))
	habitacion := ctx.Param("habitacion")

	create.HotelId = idH
	create.InitialDate = inicio
	create.FinalDate = final
	create.UserId = idU
	create.Habitacion = habitacion

	se.ReservationService.NewReserva(create)

}
