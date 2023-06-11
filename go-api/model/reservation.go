package model

import "time"

type Reservation struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"column:initial_date;not null"`
	FinalDate   time.Time `gorm:"column:final_date;not null"`
	User        User      `gorm:"foreignKey:UserId"`
	UserId      int
	Hotel       Hotel `gorm:"foreignKey:HotelId"`
}

type Reservations []Reservation
