package amenitie

import (
	amenities_dto "go-api/dto/amenitie_dto"
	"go-api/errors"
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
		log.Error(err)
		apiErr := errors.NewInternalServerErrorApi("Error al obtener amenitie por ID", err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, amenitieDto)
}

func GetAmenities(c *gin.Context) {
	var amenitiesDto amenities_dto.AmenitiesDto
	amenitiesDto, err := service.AmenitiesService.GetAmenities()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, amenitiesDto)
}

func InsertAmenitie(c *gin.Context) {
	var amenitieDto amenities_dto.AmenitieDto
	err := c.BindJSON(&amenitieDto)

	if err != nil {
		log.Error(err.Error())
		apiErr := errors.NewBadRequestErrorApi("Error al procesar la solicitud")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	amenitieDto, er := service.AmenitiesService.InsertAmenitie(amenitieDto)
	if er != nil {
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusCreated, amenitieDto)
}
