package model

type Image struct {
	Id      int    `gorm:"primaryKey"`
	Path    string `gorm:"not null"`
	HotelID int    `gorm:"not null"`
	Hotel   Hotel  `gorm:"foreignKey:HotelID"`
}

type Images []Image
