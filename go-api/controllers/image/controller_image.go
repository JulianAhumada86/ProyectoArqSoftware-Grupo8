package controllers

import (
	"go-api/dto/image_dto"
	se "go-api/services"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func InsertImage(ctx *gin.Context) {

	contentType := ctx.GetHeader("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser image/jpeg o image/png"})
		return
	}

	imagenBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los bytes de la imagen"})
		return
	}

	var img image_dto.ImageDto
	img.Data = imagenBytes
	hotelId, err := strconv.Atoi(ctx.Param("idHotel"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falta id"})
		return
	}
	img.HotelId = hotelId

	img, err = se.ImageService.InsertImage(img)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", img.Data)

} //TOKEN ADMIN

func GetImagesByHotelId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("idHotel"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parametro invalido: ID no int"})
		log.Error(err)
		return
	}

	imagesDto, err := se.ImageService.GetImagesByHotelId(id)

	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Status Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, imagesDto)

}
