package hotel

import (
	"go-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelbyid(id int) model.Hotel {
	var hotel model.Hotel
	Db.Preload("Amenities").Preload("Habitaciones").Where("id = ?", id).First(&hotel)
	log.Debug("hotel:", hotel)
	return hotel
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Preload("Amenities").Preload("Habitaciones").Find(&hotels)
	//log.Debug("Hotels: ", hotels)
	return hotels
}

func GetHabitacionesByHotelId(hotelId int) model.Amenities {
	var amenitie model.Amenities
	Db.Where("id = ?", hotelId).Find(&amenitie)
	log.Debug("amenitie:", amenitie)
	return amenitie
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

func AmenitieInHotel(am model.Hotel_amenitie) model.Hotel_amenitie {
	result := Db.Create(am)
	if result.Error != nil {
		log.Error("error")
	}

	return am
}

func HabitacionInHotel(hh model.Hotel_habitaciones) model.Hotel_habitaciones {
	result := Db.Create(hh)
	if result.Error != nil {
		log.Error("error")
	}

	return hh
}

func CantHabitaciones(idH int, idA int) int {
	var x model.Hotel_habitaciones
	Db.Where("hotel_id = ? and habitacion_id = ?", idH, idA).First(&x)

	return x.Cantidad
}
