package users_dto

type UserRequestDto struct {
	Id       int    `gorm:"prmaryKey"`
	Name     string `gorm:"type:varchar(40);not null"`
	LastName string `gorm:"type:varchar(80);not null"`
	DNI      string `gorm:"type:varchar(10);not null;unique"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
}

type UsersRequestDto struct {
	Users []UserDto `json:"users"`
}
