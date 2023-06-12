package amenitie

import (
	amenities_dto "go-api/dto/amenitie_dto"
	service "go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie ID to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var amenitieDto amenities_dto.AmenitieDto

	amenitieDto, err := service.AmenitiesService.GetAmenitiesbyid(id)

	if err != nil {
		log.Error("ERROR; FIJENSE")
		return
	}
	c.JSON(http.StatusOK, amenitieDto)
}

func GetAmenities(c *gin.Context) {
	var amenitiesDto amenities_dto.AmenitiesDto
	amenitiesDto, err := service.AmenitiesService.GetAmenities()

	if err != nil {
		log.Error("ERROR; FIJENSE")
		return
	}

	c.JSON(http.StatusOK, amenitiesDto)
}

func InsertAmenitie(ctx *gin.Context) {

	var amenitieDto amenities_dto.AmenitieDto
	nombre := ctx.Param("/name")
	amenitieDto.Name = nombre

	amenitieDto, er := service.AmenitiesService.InsertAmenitie(amenitieDto)
	if er != nil {
		log.Error("No anda, no se porque")
	}
	ctx.JSON(http.StatusCreated, amenitieDto)
}
