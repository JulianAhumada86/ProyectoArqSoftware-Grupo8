package user

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	userdto "go-api/dto/users_dto"
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

// name/:LastName/:DNI/:Password/:Email/:Admin
func AddUser(ctx *gin.Context) {
	var userDto userdto.UserDto

	userDto.Name = ctx.Param(";name")
	userDto.LastName = ctx.Param(":LastName")
	userDto.DNI = ctx.Param(":DNI")
	userDto.Password = ctx.Param(":Password")
	userDto.Email = ctx.Param(":Email")
	ad, err := strconv.Atoi(ctx.Param(":Admin"))
	userDto.Admin = ad

	if err != nil {
		log.Error("Algo falla y no me importa que")
		return
	}

	userDto, err = se.UserService.AddUser(userDto)

	if err != nil {
		log.Error("Algo falla y no me importa que")
		return
	}

	ctx.JSON(http.StatusOK, userDto)
}
