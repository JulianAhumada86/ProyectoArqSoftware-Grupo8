package services

import (
	uClient "go-api/clients/user"
	"go-api/dto/users_dto"
	uDto "go-api/dto/users_dto"
	e "go-api/errors"
	"go-api/model"
	"go-api/users_dto"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (uDto.UserDto, error)
	GetUsers() (uDto.UserDto, e.ErrorApi)
	AddUser(userDto uDto.UserDto) (uDto.UserDto, error)
	Login(loginDto users_dto.LoginDto) (users_dto.LoginResponseDto, e.ErrorApi)
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

func (s *userService) AddUser(userDto uDto.UserDto) (uDto.UserDto, e.ErrorApi) {
	var userModel model.User

	userModel.Admin = userDto.Admin
	userModel.Email = userDto.Email
	userModel.DNI = userDto.DNI
	userModel.LastName = userDto.LastName
	userModel.Password = userDto.Password
	userModel.Name = userDto.Name

	uClient.AddUser(userModel)
	userDto.Id = userModel.Id
	return userDto, nil
}

func (s *userService) GetUsers() (uDto.UserDto, e.ErrorApi) {
	var users model.Users = uClient.GetUsers()
	uDto := users_dto.UsersDto{
		Users: make([]users_dto.UserDto, len(users)),
	}

	for i, user := range users {
		userDto := users_dto.UserDto{
			Id:       user.Id,
			Name:     user.Name,
			LastName: user.LastName,
			DNI:      user.DNI,
			Email:    user.Email,
			Admin:    user.Admin,
		}

		uDto.Users[i] = userDto
	}
	return uDto, nil //aca no se que tengo que cambiar para que este bien
}

func (s *userService) Login(loginDto users_dto.LoginDto) (users_dto.LoginResponseDto, e.ErrorApi) {
	var user model.User
	user, err := uClient.GetUserByUsername(loginDto.Username)
	var loginResponseDto users_dto.LoginResponseDto
	loginResponseDto.UserId = -1

	if err != nil {
		return loginResponseDto, e.NewBadRequestErrorApi("Usuario no encontrado")
	}

	var comparison error = s.VerifyPassword(user.Password, loginDto.Password)

	if loginDto.Username == user.UserName {
		{
			if comparison != nil {
				return loginResponseDto, e.NewUnauthorizedErrorApi("contraseña incorrecta")
			}
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"password": loginDto.Password,
	})
	var jwtKey = []byte("Secret key")
	tokenString, _ := token.SignedString(jwtKey)

	var verifyToken error = s.VerifyPassword(user.Password, tokenString)

	if loginDto.Username != user.UserName {
		if verifyToken != nil {
			return loginResponseDto, e.NewUnauthorizedErrorApi("Contraseña incorrecta")
		}
	}

	loginResponseDto.UserId = user.Id
	loginResponseDto.Token = tokenString
	loginResponseDto.Type = user.Type
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

//Esto es para  AddUser
/*func (s *userService) HashPassword(password string)(string, error){
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("No fue posible generar un Hash a partir de la password %w", err)
	}

	return string(HashedPassword), nil
}
*/
