package amenitie

import (
	"go-api/errors"
	//service "go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie ID to load: " + c.Param("id"))

	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg := "Error al convertir ID a entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	//_, err = service.AmenitiesService.GetAmenitiesbyid(id)

	if err != nil {
		log.Error(err)
		apiErr := errors.NewInternalServerErrorApi("Error al obtener amenitie por ID", err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, err)
}

func GetAmenities(c *gin.Context) {
	//amenitiesDto, err := service.AmenitiesService.GetAmenities()
	err := c.Param("hola")
	/*
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	*/
	c.JSON(http.StatusOK, err)
}

func InsertAmenitie(c *gin.Context) {
	err := c.Param("hola")
	/*
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
	*/
	c.JSON(http.StatusCreated, err)
}
