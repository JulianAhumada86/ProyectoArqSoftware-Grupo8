package user

import (
	"go-api/dto/users_dto"
	userdto "go-api/dto/users_dto"
	e "go-api/errors"
	se "go-api/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// Controlador para obtener un usuario por su ID
func GetUserById(ctx *gin.Context) {
	log.Debug("User id to load: " + ctx.Param("id"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	var userDto users_dto.UserDto
	userDto, err := se.UserService.GetUserById(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, userDto)

	/*
		controllers.VerificacionToken()(ctx)
		if err := ctx.Errors.Last(); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		userID := ctx.GetInt("user_id")
		if !controllers.Admin(userID){
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Debe tener permiso de administrador para realizar esta accion"})
			return
		}

		log.Debug("User if to load: " + ctx.Param("id"))

		id, _ := strconv.Atoi(ctx.Param("id"))
		var userDto userdto.UserDto

		userDto, err := se.UserService.GetUserById(id)

		if err != nil {
			ctx.JSON(err.Status(), err) //aca err status hay que implementarlo en manejo de errores
			return
		}

		ctx.JSON(http.StatusOK, userDto)
	*/
	/*
		id, _ := strconv.Atoi(ctx.Param("id"))
		userDto, err := se.UserService.GetUserById(id)

		if err != nil {
			log.Print("error")
		}
		ctx.JSON(http.StatusOK, userDto)
	*/
}

func GetUsers(ctx *gin.Context) {
	var userDto users_dto.UserDto
	userDto, err := se.UserService.GetUsers()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, userDto)
}

// name/:LastName/:DNI/:Password/:Email/:Admin
func AddUser(ctx *gin.Context) {

	/*

		err := ctx.BindJSON(&userDto)

		// Error Parsing json param
		if err != nil {
			log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userDto, e := se.UserService.AddUser(userDto)
		// Error del Insert
		if e != nil {
			//ctx.JSON(e.Status(), err) //no puedo conectarlo al archivo errors.go
			return
		}

		ctx.JSON(http.StatusCreated, userDto)
	*/
	var userDto userdto.UserDto

	userDto.Name = ctx.Param(":name")         //
	userDto.LastName = ctx.Param(":LastName") //
	userDto.DNI = ctx.Param(":DNI")           //
	userDto.Password = ctx.Param(":Password") //
	userDto.Email = ctx.Param(":Email")       //
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

func Login(ctx *gin.Context) {
	var loginDto users_dto.LoginDto
	err := ctx.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Debug(loginDto)
	e.Status()
	//var loginResponseDto users_dto.LoginResponseDto
	loginResponseDto, err := se.UserService.Login(loginDto)
	if err != nil {
		if err.Status() == 400 {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, loginResponseDto)

}
