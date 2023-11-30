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

func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)
	log.Debug("Hotels: ", hotels)
	return hotels
}

func InsertHotel(hotel model.Hotel) model.Hotel {
	h := Db.Create(&hotel)

	if h.Error != nil {
		log.Error("")
	}

	log.Debug("hotel creado. Id= ", hotel.Id)

	return hotel

}

func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie
	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("amenitie:", amenitie)
	return amenitie
}

/*modifique el insert*/
func InsertAmenitie(amenitie_model model.Amenitie) model.Amenitie {
	result := Db.Create(amenitie_model)
	if result.Error != nil {
		log.Error("error")
	}
	log.Debug("Amenitie creada: ", amenitie_model.Id)
	return amenitie_model
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities

	Db.Find(&amenities)
	log.Debug("Amenities: ", amenities)

	return amenities
}
