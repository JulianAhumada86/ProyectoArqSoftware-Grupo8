package model

import "time"

type Reservation struct {
	Id           int       `gorm:"primaryKey"`
	InitialDate  time.Time `gorm:"column:initial_date;not null"`
	FinalDate    time.Time `gorm:"column:final_date;not null"`
	HabitacionId int
	UserID       int
	HotelID      int
	User         User       `gorm:"foreignKey:UserID"`
	Hotel        Hotel      `gorm:"foreignKey:HotelID"`
	Habitacion   Habitacion `gorm:"foreignKey:HabitacionId"`
}

type Reservations []Reservation
