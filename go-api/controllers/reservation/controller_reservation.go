package reservation

import (
	"go-api/dto/reservations_dto"
	reservationDTO "go-api/dto/reservations_dto"
	"go-api/errors"
	se "go-api/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func NewReserva(ctx *gin.Context) {
	var create reservationDTO.ReservationCreateDto

	idH, err := strconv.Atoi(ctx.Param("idHotel"))
	if err != nil {
		errMsg := "Error al convertir ID de Hotel a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	inicio := ctx.Param("inicio")
	final := ctx.Param("final")

	idU, err := strconv.Atoi(ctx.Param("idUser"))
	if err != nil {
		errMsg := "Error al convertir ID de Usuario a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	habitacion, err := strconv.Atoi(ctx.Param("habitacion"))
	if err != nil {
		errMsg := "Error al convertir ID de Habitación a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	create.HotelId = idH
	create.InitialDate = inicio
	create.FinalDate = final
	create.UserId = idU
	create.HabitacionId = habitacion

	reservationDTO, err := se.ReservationService.NewReserva(create)

	if err != nil {
		log.Error(err)
		apiErr := errors.NewInternalServerErrorApi("Error al crear nueva reserva", err)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	ctx.JSON(http.StatusCreated, reservationDTO)
}

//TOKEN Cliente

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

func Disponibilidad_de_reserva(ctx *gin.Context) {
	var reserva reservationDTO.ReservationCreateDto

	idH, err := strconv.Atoi(ctx.Param("idHotel"))
	if err != nil {
		errMsg := "Error al convertir ID de Hotel a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	inicio := ctx.Param("inicio")
	final := ctx.Param("final")

	idU, err := strconv.Atoi(ctx.Param("idUser"))
	if err != nil {
		errMsg := "Error al convertir ID de Usuario a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	habitacion, err := strconv.Atoi(ctx.Param("habitacion"))
	if err != nil {
		errMsg := "Error al convertir ID de Habitación a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	reserva.HotelId = idH
	reserva.InitialDate = inicio
	reserva.FinalDate = final
	reserva.UserId = idU
	reserva.HabitacionId = habitacion

	err = se.ReservationService.Disponibilidad_de_reserva(reserva)

	if err == nil {
		ctx.JSON(http.StatusOK, reserva)
	} else {
		log.Error(err)
		apiErr := errors.NewBadRequestErrorApi("Error al verificar disponibilidad de reserva")
		ctx.JSON(apiErr.Status(), apiErr)
	}
}

func GetReservasByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		errMsg := "Error al convertir ID de Usuario a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	var reservasDto reservations_dto.ReservationsDto
	reservasDto, err = se.ReservationService.GetReservasByUserId(id)

	if err != nil {
		log.Error(err)
		apiErr := errors.NewInternalServerErrorApi("Error al obtener reservas por ID de Usuario", err)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	ctx.JSON(http.StatusOK, reservasDto)
}

//Token Client
