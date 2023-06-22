package reservation

import (
	"go-api/dto/reservations_dto"
	reservationDTO "go-api/dto/reservations_dto"
	se "go-api/services"
	"net/http"
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

	reservationDTO, err := se.ReservationService.NewReserva(create)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusCreated, reservationDTO)

}

func GetReservaById(ctx *gin.Context) {
	var create reservationDTO.ReservationDto
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: ID no int"})
	}
	create = se.ReservationService.GetReservaById(id)

	ctx.JSON(http.StatusOK, create)
}

func GetReservas(ctx *gin.Context) {
	var reservasDto reservations_dto.ReservationsDto
	reservasDto, err := se.ReservationService.GetReservas()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
}

/* FUTURA FUNCION
func GetReservasByUserId(ctx *gin.Context) {
	log.Debug("user id to load: " + ctx.Param("user_id"))
	id, _ := strconv.Atoi(ctx.Param("user_id"))
	var reservasDto reservations_dto.ReservationsDto
	reservasDto, err := se.ReservationService.GetReservaByUserId(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
}
*/
