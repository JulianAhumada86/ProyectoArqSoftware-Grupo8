package model

type Hotel_habitaciones struct {
	HabitacionID int `gorm:"primaryKey"`
	HotelID      int `gorm:"primaryKey"`
	Cantidad     int
}
