package hotel

// Funciones que hay que poner
// get y post

import (
	"log"
	"net/http"
	"strconv"

	hotel_dto "go-api/dto/hotels_dto"
	se "go-api/services"

	"github.com/gin-gonic/gin"
)

func GetHotelbyid(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var hotelDto hotel_dto.HotelDto

	hotelDto, err := se.HotelService.GetHotelbyid(id)
	ctx.JSON(http.StatusOK, hotelDto)

	if err != nil {
		log.Print("error")
	}
}

func InsertHotel(ctx *gin.Context) {
	//Aca hay que hacer el token verificacion, pero in pta idea como es, lo vemos despues.

	var hotelDto hotel_dto.HotelDto

	hotelDto.Name = ctx.Param("name")
	num, _ := strconv.Atoi(ctx.Param("Nroom"))
	hotelDto.RoomsAvailable = num
	hotelDto.Description = ctx.Param("descr")

	hotelDto, err := se.HotelService.InsertHotel(hotelDto)

	if err != nil {
		log.Println("Error")
	}
	ctx.JSON(http.StatusOK, hotelDto)

}
