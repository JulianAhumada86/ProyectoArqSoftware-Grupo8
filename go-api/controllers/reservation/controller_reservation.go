package reservation

import (
	"go-api/dto/reservations_dto"
	reservationDTO "go-api/dto/reservations_dto"
	"go-api/errors"
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
	habitacion, _ := strconv.Atoi(ctx.Param("habitacion"))

	create.HotelId = idH
	create.InitialDate = inicio
	create.FinalDate = final
	create.UserId = idU
	create.HabitacionId = habitacion

	reservationDTO, err := se.ReservationService.NewReserva(create)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, reservationDTO)

} //TOKEN Cliente

func GetReservaById(ctx *gin.Context) {
	var create reservationDTO.ReservationDto
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		errMsg := "ID inválido"
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}
	create = se.ReservationService.GetReservaById(id)

	ctx.JSON(http.StatusOK, create)
} //Token Admin

func GetReservas(ctx *gin.Context) {
	var reservasDto reservations_dto.ReservationsDto
	reservasDto, err := se.ReservationService.GetReservas()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
} //TOken Client

func Dispoibilidad_de_reserva(ctx *gin.Context) {
	var reserva reservationDTO.ReservationCreateDto

	idH, err := strconv.Atoi(ctx.Param("idHotel"))
	inicio := ctx.Param("inicio")
	final := ctx.Param("final")

	idU, _ := strconv.Atoi(ctx.Param("idUser"))
	habitacion, _ := strconv.Atoi(ctx.Param("habitacion"))

	reserva.HotelId = idH
	reserva.InitialDate = inicio
	reserva.FinalDate = final
	reserva.UserId = idU
	reserva.HabitacionId = habitacion
	err = se.ReservationService.Disponibilidad_de_reserva(reserva)
	if err == nil {
		ctx.JSON(http.StatusOK, reserva)
	} else {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

}

func GetReservasByUserId(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("user_id"))

	var reservasDto reservations_dto.ReservationsDto
	reservasDto, err := se.ReservationService.GetReservasByUserId(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
} //Token Client
