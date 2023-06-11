package hotel

import (
	"go-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelbyid(id int) model.Hotel {
	var hotel model.Hotel
	Db.Where("id = ?", id).First(&hotel)
	log.Debug("hotel:", hotel)
	return hotel

}

func InsertHotel(hotel model.Hotel) model.Hotel {
	h := Db.Create(&hotel)

	if h.Error != nil {
		log.Error("")
	}

	log.Debug("hotel creado. Id= ", hotel.Id)

	return hotel

}
