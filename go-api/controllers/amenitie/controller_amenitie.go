package amenitie

import (
	amenities_dto "go-api/dto/amenitie_dto"
	service "go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) { //Verificar Token admin
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

/*Modifique el Insert*/
func InsertAmenitie(c *gin.Context) {
	var amenitieDto amenities_dto.AmenitieDto
	err := c.BindJSON(&amenitieDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitieDto, er := service.AmenitiesService.InsertAmenitie(amenitieDto)
	if er != nil {
		log.Error("no se como implementar lo del paquete errors")
		return
	}

	c.JSON(http.StatusCreated, amenitieDto)
}
