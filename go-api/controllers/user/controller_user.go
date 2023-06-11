package controllers

import (
	"go-api/controllers"
	"net/http"
	"strconv"

	user_dto "go-api/dto/users_dto"
	service "go-api/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Controlador para obtener un usuario por su ID
func GetUserById(c *gin.Context) {
	controllers.VerificacionToken()(c)
	//si hubo un error en la verificacion del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//obtenemos el ID del contexto
	userID := c.GetInt("user_id")

	//verificamos si es admin o no
	if !controllers.Admin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Necesita permiso de administrador"})
		return
	}

	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var userDto users_dto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

// controladores para obtener los usuarios
func GetUsers(c *gin.Context) {
	controllers.VerificacionToken()(c)
	//si hubo un error en la verificacion del token
	if err := c.Errors.Last(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//obtenemos el ID del contexto
	userID := c.GetInt("user_id")

	//verificamos si es admin o no
	if !controllers.Admin(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Necesita permiso de administrador"})
		return
	}

	var usersDto user_dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}
