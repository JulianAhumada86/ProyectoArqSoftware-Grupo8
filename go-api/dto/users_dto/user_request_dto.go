package users_dto

type UserRequestDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	DNI      string `json:"dni"`
	Email    string `json:"email"`
	Admin    int    `json:"admin"`
	Token    string `json:"token"`
}

type UsersRequestDto struct {
	Users []UserDto `json:"users"`
}
