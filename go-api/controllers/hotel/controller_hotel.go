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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: ID no int"})
		log.Error(err)
		return
	}
	var hotelDto hotel_dto.HotelDto

	hotelDto, err = se.HotelService.GetHotelbyid(id)

	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Status Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, hotelDto)
}

func InsertHotel(ctx *gin.Context) {

	var hotelDto hotel_dto.HotelDto

	hotelDto.Name = ctx.Param("name")
	num, err := strconv.Atoi(ctx.Param("Nroom"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: Nroom no int"})
		return
	}
	hotelDto.RoomsAvailable = num
	hotelDto.Description = ctx.Param("descr")

	if hotelDto.Name == "" {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: Argumento vacio/no name"})
		return
	}
	if hotelDto.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: Argumento vacio/no descrpiption"})
		return
	}

	hotelDto, err = se.HotelService.InsertHotel(hotelDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, hotelDto)

} //ADMIN Token
