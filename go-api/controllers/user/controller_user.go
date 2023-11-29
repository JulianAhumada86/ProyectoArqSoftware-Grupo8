package user

import (
	"go-api/dto/users_dto"
	userdto "go-api/dto/users_dto"
	se "go-api/services"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
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

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    userDto.Email,
			"password": userDto.Password,
			"admin":    userDto.Admin,
		})

		tokenString, err := token.SignedString([]byte("Secret key"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
			return
		}
		userDto.Token = tokenString

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
	log.Println("Token es:")
	log.Print(respuesta.Token)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    loginDto.Email,
		"password": loginDto.Password,
		"admin":    respuesta.Admin,
	})

	tokenString, err := token.SignedString([]byte("Secret key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		return
	}
	respuesta.Token = tokenString
	log.Println(tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err) //Mandar error
		return
	}

	log.Println("Token es:")
	log.Print(respuesta.Token)
	ctx.JSON(http.StatusOK, respuesta)
}
