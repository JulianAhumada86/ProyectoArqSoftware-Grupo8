package image

import (
	"go-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertImage(im model.Image) (model.Image, error) {
	img := Db.Create(&im)

	if img.Error != nil {
		return im, img.Error
	}

	log.Debug("imagen para hotel. Id= ", im.HotelID)

	return im, nil
}

func GetImagesByHotelId(id int) (model.Images, error) {
	var imgs model.Images
	Db.Where("hotel_id = ?", id).Find(&imgs)
	if Db.Error != nil {
		return nil, Db.Error
	}

	return imgs, nil
}
