package user

import (
	"go-api/dto/users_dto"
	userdto "go-api/dto/users_dto"
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
	//var userDto users_dto.UsersDto
	userDto, err := se.UserService.GetUsers()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, userDto)
}

// name/:LastName/:DNI/:Password/:Email/:Admin
func AddUser(ctx *gin.Context) {
	var userDto userdto.UserDto
	var userRdto userdto.UserRequestDto
	userDto.Name = ctx.Param("name")
	if userDto.Name == "" {
		log.Error("El campo 'name' está vacío")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	userDto.LastName = ctx.Param("LastName")
	if userDto.LastName == "" {
		log.Error("El campo 'LastName' está vacío")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "LastName is required"})
		return
	}
	userDto.DNI = ctx.Param("DNI")
	if userDto.DNI == "" {
		log.Error("El campo 'DNI' está vacío")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "DNI is required"})
		return
	}
	userDto.Password = ctx.Param("Password")
	if userDto.Password == "" {
		log.Error("El campo 'Password' está vacío")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}
	userDto.Email = ctx.Param("Email")
	if userDto.Email == "" {
		log.Error("El campo 'Email' está vacío")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	ad, err := strconv.Atoi(ctx.Param("Admin"))
	userDto.Admin = ad

	if err != nil {
		log.Error("Algo falla y no me importa que")
		return
	}

	userRdto, err = se.UserService.AddUser(userDto)

	if err != nil {
		log.Error("Algo falla al llamar al service para agregar el usuario")
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, userRdto)
	}

}

func Login(ctx *gin.Context) {

	var loginDto users_dto.LoginDto
	err := ctx.BindJSON(&loginDto)
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	respuesta, err := se.UserService.Login(loginDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respuesta)
		return
	}

	respuesta.DNI = "44193211"
	respuesta.LastName = "Ahumada"
	respuesta.Name = "Julian"
	respuesta.Email = "2214335@ucc.edu.ar"
	ctx.JSON(http.StatusOK, respuesta)
}
