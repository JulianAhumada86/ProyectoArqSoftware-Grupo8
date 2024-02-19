package model

type Habitacion struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"varchar(40);not null"`
	Camas  int
	Piezas int      `gorm:"int;not null"`
	Hotels []*Hotel `gorm:"many2many:hotel_habitaciones"`
}

type Habitaciones []Habitacion
