package model

type Image struct {
	Id      int    `gorm:"primaryKey"`
	Imagen  []byte `gorm:"not null"`
	HotelID int    `gorm:"not null"`
	Hotel   Hotel  `gorm:"foreignKey:HotelID"`
}

type Images []Image
