package users_dto

type UserDto struct {
	Id       int    `gorm:"prmaryKey"`
	Name     string `gorm:"type:varchar(40);not null"`
	LastName string `gorm:"type:varchar(80);not null"`
	UserName string `json:"username"`
	DNI      string `gorm:"type:varchar(10);not null;unique"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(30);not null"`
	Type     bool   `json:"type"`
	Admin    int
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}
