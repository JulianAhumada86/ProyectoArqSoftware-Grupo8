package hotel

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	hotel_dto "go-api/dto/hotels_dto"
	se "go-api/services"

	"github.com/gin-gonic/gin"
)

func GetHotelbyid(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	var hotelDto hotel_dto.HotelDto

	hotelDto, err = se.HotelService.GetHotelbyid(id)

	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, hotelDto)
}

func GetHotels(ctx *gin.Context) {
	var hotelesDto hotel_dto.HotelsDto
	hotelesDto, err := se.HotelService.GetHotels()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	log.Println(hotelesDto)
	ctx.JSON(http.StatusOK, hotelesDto)
} //TOken Client

func GetHotelsC(ctx *gin.Context) {
	var hotelesDto hotel_dto.HotelsDto
	hotelesDto, err := se.HotelService.GetHotelsC()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, hotelesDto)
} //TOken Client

func InsertHotel(ctx *gin.Context) {

	var hotelDto hotel_dto.HotelConHabitaciones
	err := ctx.ShouldBindJSON(&hotelDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, err = se.HotelService.InsertHotel(hotelDto)

	ctx.JSON(http.StatusOK, hotelDto)

} //ADMIN Token
