package services

import (
	uClient "go-api/clients/user"
	uDto "go-api/dto/users_dto"
	"go-api/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (uDto.UserDto, error)
	AddUser(userDto uDto.UserDto) (uDto.UserDto, error)
}

var UserService userServiceInterface

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (uDto.UserDto, error) {
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

func (s *userService) AddUser(userDto uDto.UserDto) (uDto.UserDto, error) {
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
