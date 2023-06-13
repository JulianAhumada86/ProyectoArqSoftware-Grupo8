package model

import "time"

type Reservation struct {
	Id          int       `gorm:"primaryKey"`
	InitialDate time.Time `gorm:"column:initial_date;not null"`
	FinalDate   time.Time `gorm:"column:final_date;not null"`
	Habitacion  string
	UserID      int
	HotelID     int
	User        User   `gorm:"foreignKey:UserID"`
	Hotel       Hotel  `gorm:"foreignKey:HotelID"`
	Address     string `json:"hotel_address"`
}

type Reservations []Reservation
