package users_dto

type UserDto struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Amenities      []string `json:"amenities"`
	RoomsAvailable int      `json:"rooms_available"`
}

type UsersDto struct {
	Hotels []UserDto `json:"hotels"`
}
