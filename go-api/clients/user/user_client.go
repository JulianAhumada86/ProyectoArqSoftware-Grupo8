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

func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := Db.Where("user_name = ?", username).First(&user)

	log.Debug("User: ", user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) bool {
	var user model.User
	result := Db.Where("email = ?", email).First(&user)

	log.Debug("User: ", user)

	if result.Error != nil {
		return true //si no lo encuentra --> no existe
	}

	return false
}

func AddUser(user model.User) model.User {
	Db.Create(&user)
	log.Debug("User:", user)
	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}
