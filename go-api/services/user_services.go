package services

import (
	"fmt"
	uClient "go-api/clients/user"
	"go-api/dto/users_dto"
	uDto "go-api/dto/users_dto"
	e "go-api/errors"
	"go-api/model"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (uDto.UserDto, e.ErrorApi)
	GetUsers() (uDto.UsersDto, e.ErrorApi)
	AddUser(userDto uDto.UserDto) (uDto.UserRequestDto, e.ErrorApi)
	Login(loginDto users_dto.LoginDto) (users_dto.UserRequestDto, e.ErrorApi)

	HashPassword(string) (string, error)
	VerifyPassword(string, string) error
}

var UserService userServiceInterface

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (uDto.UserDto, e.ErrorApi) {
	userModel := uClient.GetUserById(id)

	var userDto uDto.UserDto
	userDto.Admin = userModel.Admin
	userDto.Email = userModel.Email
	userDto.DNI = userModel.DNI
	userDto.LastName = userModel.LastName
	userDto.Password = userModel.Password
	userDto.Name = userModel.Name
	userDto.Id = id

	return userDto, nil
}

func (s *userService) AddUser(userDto uDto.UserDto) (uDto.UserRequestDto, e.ErrorApi) {
	var userModel model.User
	var userRDto uDto.UserRequestDto

	userModel.Admin = userDto.Admin
	userModel.Email = userDto.Email
	userModel.DNI = userDto.DNI
	userModel.LastName = userDto.LastName

	hashedPassword, _ := UserService.HashPassword(userDto.Password)

	userModel.Password = hashedPassword
	userModel.Name = userDto.Name

	if uClient.ExistUserByDni(userModel.DNI) {
		log.Error("Algo salio mal en el email")
		return userRDto, e.NewBadRequestErrorApi("DNI Ya Registrado")
	}
	if uClient.ExistUserByEmail(userModel.Email) {
		log.Error("Algo salio mal en el email")
		return userRDto, e.NewBadRequestErrorApi("Email Ya registrado")
	}

	uClient.AddUser(userModel)

	userRDto.Name = userDto.Name
	userRDto.LastName = userDto.LastName
	userRDto.Email = userDto.Email
	userRDto.DNI = userDto.DNI
	userRDto.Id = userModel.Id
	return userRDto, nil

}

func (s *userService) GetUsers() (uDto.UsersDto, e.ErrorApi) {
	var users model.Users = uClient.GetUsers()

	usersList := make([]uDto.UserDto, 0)
	for _, user := range users {
		var userDto users_dto.UserDto //de cada uno
		if !userDto.Type {
			userDto.Name = user.Name
			userDto.LastName = user.LastName

			userDto.Email = user.Email
			userDto.Id = user.Id
			userDto.Type = user.Type
		}

		usersList = append(usersList, userDto) //ver esto
	}
	return uDto.UsersDto{
		Users: usersList,
	}, nil

}

func (s *userService) Login(loginDto users_dto.LoginDto) (users_dto.UserRequestDto, e.ErrorApi) {
	var userRequestDto users_dto.UserRequestDto
	userRequestDto.Email = loginDto.Email
	user, err := uClient.GetUserByEmail(loginDto.Email)

	if err != nil {
		return userRequestDto, e.NewBadRequestErrorApi("Email no existe o no esta registrado")
	}

	var comparison error = s.VerifyPassword(user.Password, loginDto.Password)

	if loginDto.Email == user.Email {
		if comparison != nil {
			log.Error(comparison)
			return userRequestDto, e.NewUnauthorizedErrorApi("contraseña incorrecta")
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    loginDto.Email,
		"password": loginDto.Password,
	})

	var jwtKey = []byte("Secret key")
	tokenString, _ := token.SignedString(jwtKey)

	var verifyToken error = s.VerifyPassword(user.Password, tokenString)

	if loginDto.Email != user.Email {
		if verifyToken != nil {
			return userRequestDto, e.NewUnauthorizedErrorApi("Contraseña incorrecta")
		}
	}

	userRequestDto.Name = user.Name
	userRequestDto.LastName = user.LastName
	userRequestDto.DNI = user.DNI
	userRequestDto.Id = user.Id
	return userRequestDto, nil
}

func (s *userService) VerifyPassword(HashedPassword string, candidatePassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(candidatePassword))
}

func (s *userService) HashPassword(password string) (string, error) {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	log.Println("EL largo del ash es: " + string(len(string(HashedPassword))))

	if err != nil {
		return "", fmt.Errorf("No fue posible generar un Hash a partir de la password %w", err)
	}

	return string(HashedPassword), nil
}
