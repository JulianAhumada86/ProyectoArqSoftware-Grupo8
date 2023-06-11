package controllers

import (
	"go-api/controllers"
	"net/http"
	"strconv"

	user_dto "go-api/dto/user_dto"
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

// controlador para insertar usuario nuevo
func UserInsert(c *gin.Context) {
	var userDto user_dto
	err := c.BindJSON(&userDto)

	//error parsinfg json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//verificamos que no haya campos vacios
	if controllers.IsEmptyField(userDto.name) || controllers.IsEmptyField(userDto.LastName) ||
		controllers.IsEmptyField(userDto.Mail) || controllers.IsEmptyField(userDto.Dni) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios estan vacios"})
		return
	}

	// Verificamos si ya existe el mail
	if service.UserService.IsEmailTaken(userDto.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El Email ingresado ya se encuentra registrado"})
		return
	}

	var userDetailDto user_dto
	userDetailDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDetailDto)

}

// controlador para el inicio de sesión de un usuario.
func UserLogin(c *gin.Context) {
	var userDto user_dto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Verificar si alguno de los campos está vacío
	if controllers.IsEmptyField(userDto.Email) ||
		controllers.IsEmptyField(userDto.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uno o varios de los campos obligatorios esta vacio o no se envio"})
		return
	}

	var loginResponse user_dto.UserLoginResponseDto
	loginResponse, er := service.UserService.UserLogin(userDto)
	// Error del Login
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}
