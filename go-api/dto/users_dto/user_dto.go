package users_dto

// Sacar los gorm po json
type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	DNI      string `json:"dni"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     bool   `json:"type"`
	Admin    int    `json:"admin"`
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}
