package controllers

import (
	"log"
	"net/http"
	"strconv"

	se "go-api/services"

	"github.com/gin-gonic/gin"
)

// Controlador para obtener un usuario por su ID
func GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userDto, err := se.UserService.GetUserById(id)

	if err != nil {
		log.Print("error")
	}
	ctx.JSON(http.StatusOK, userDto)

}
