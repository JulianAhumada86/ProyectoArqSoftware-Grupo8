package model

type Image struct {
	Id      int    `gorm:"primaryKey"`
	Imagen  []byte `gorm:"type:longblob;not null"`
	HotelID int    `gorm:"not null"`
	Hotel   Hotel  `gorm:"foreignKey:HotelID"`
}

type Images []Image
