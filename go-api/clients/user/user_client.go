package user

import (
	"go-api/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User:", user)

	return user
}

func AddUser(user model.User) model.User {
	Db.Create(&user)
	log.Debug("User:", user)

	return user
}
