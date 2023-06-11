package user_dto

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"las_name"`
	Mail     string `json:"mail"`
	Dni      string `json:"dni"`
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}
