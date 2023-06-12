package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(40);not null"`
	LastName string `gorm:"type:varchar(80);not null"`
	UserName string `gorm:"type:varchar(200);not null;unique"`
	DNI      string `gorm:"type:varchar(10);not null;unique"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(30);not null"`
	Type     bool   `gorm:"not null;default:false"`
	Admin    int

	Reservations Reservations `gorm:"foreignKey:UserId"`
}

type Users []User
