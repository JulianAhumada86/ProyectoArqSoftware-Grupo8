package user

import (
	"go-api/dto/users_dto"
	userdto "go-api/dto/users_dto"
	"go-api/errors"
	se "go-api/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// Controlador para obtener un usuario por su ID
func GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("idUsuario"))

	admin, err2 := strconv.Atoi(ctx.Param("admin"))

	if err != nil {
		//Si no mandaron idUsuario entra aca, por lo que el usuario se busca a si mismo
		id, err = strconv.Atoi(ctx.Param("id"))
		log.Info("Un usuario se busca a si mismo")
		if err != nil {
			errMsg := "Error al convertir ID a entero"
			log.Error(errMsg)
			apiErr := errors.NewBadRequestErrorApi(errMsg)
			ctx.JSON(apiErr.Status(), apiErr)
			return
		}
	} else if err2 != nil {
		errMsg := "Error al convertir ID del admin entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return

	} else if admin == 1 {

		log.Info("Un administrador esta buscando a un usuario")
	} else {
		errMsg := "Error al convertir ID del admin entero"
		log.Error(errMsg)
		apiErr := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	var userDto users_dto.UserDto
	userDto, err = se.UserService.GetUserById(id)

	if err != nil {
		apiErr, ok := err.(errors.ErrorApi)
		if !ok {
			errMsg := "Error interno del servidor"
			log.Error(errMsg)
			apiErr = errors.NewInternalServerErrorApi(errMsg, err)
		}
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}
	ctx.JSON(http.StatusOK, userDto)
}

func GetUsers(ctx *gin.Context) {
	//var userDto users_dto.UsersDto
	userDto, err, l := se.UserService.GetUsers()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	log.Print(l)
	ctx.Header("X-Total-Count", strconv.Itoa(l))
	ctx.JSON(http.StatusOK, userDto)
}

// name/:LastName/:DNI/:Password/:Email/:Admin
func AddUser(ctx *gin.Context) {
	var userDto userdto.UserDto
	var userRdto userdto.UserRequestDto
	userDto.Name = ctx.Param("name")
	if userDto.Name == "" {
		errMsg := "El campo 'name' está vacío"
		log.Error(errMsg)
		err := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(err.Status(), err)
		return
	}
	userDto.LastName = ctx.Param("LastName")
	if userDto.LastName == "" {
		errMsg := "El campo 'LastName' está vacío"
		log.Error(errMsg)
		err := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(err.Status(), err)
		return
	}
	userDto.Admin = 0 //Pongo 0 porque por defecto no es admin
	userDto.DNI = ctx.Param("DNI")
	if userDto.DNI == "" {
		errMsg := "El campo 'DNI' está vacío"
		log.Error(errMsg)
		err := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(err.Status(), err)
		return
	}
	userDto.Password = ctx.Param("Password")
	if userDto.Password == "" {
		errMsg := "El campo 'Password' está vacío"
		log.Error(errMsg)
		err := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(err.Status(), err)
		return
	}
	userDto.Email = ctx.Param("Email")
	if userDto.Email == "" {
		errMsg := "El campo 'Email' está vacío"
		log.Error(errMsg)
		err := errors.NewBadRequestErrorApi(errMsg)
		ctx.JSON(err.Status(), err)
		return
	}

	userRdto, err := se.UserService.AddUser(userDto)

	if err != nil {
		errMsg := "Algo falla al llamar al service para agregar el usuario"
		log.Error(errMsg)
		apiErr := errors.NewInternalServerErrorApi(errMsg, err)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	ctx.JSON(http.StatusOK, userRdto)

}

func Login(ctx *gin.Context) {

	var loginDto users_dto.LoginDto
	err := ctx.BindJSON(&loginDto)
	if err != nil {
		log.Error(err.Error())
		apiErr := errors.NewBadRequestErrorApi(err.Error())
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}
	respuesta, err := se.UserService.Login(loginDto)
	if err != nil {
		errMsg := "Credenciales de inicio de sesión inválidas"
		log.Error(errMsg)
		apiErr := errors.NewForbiddenErrorApi(errMsg)
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}
	ctx.JSON(http.StatusOK, respuesta)
}
