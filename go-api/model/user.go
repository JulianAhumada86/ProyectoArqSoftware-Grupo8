package model

type User struct {
	Id       int    `gorm:"prmaryKey"`
	Name     string `gorm:"type:varchar(40);not null"`
	LastName string `gorm:"type:varchar(80);not null"`
	DNI      string `gorm:"type:varchar(10);not null;unique"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(30);not null"`
	Admin    int
	Reserva  Reservations `gorm:"foreignKey:UserId"`
}

type Users []User
