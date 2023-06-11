package model

type Hotel struct {
	Id             int    `gorm:"primaryKey"`
	Name           string `gorm:"varchar(40);not null"`
	RoomsAvailable int
	Description    string `gorm:"varchar(400);not null"`
	//Amenities      []Amenitie `gorm:"many2many:hotel_amenities"`
	//Images         []Image    `gorm:"foreignKey:HotelId"`
}

type Hotels []Hotel
