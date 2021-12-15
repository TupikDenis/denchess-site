package user

import (
	"log"

	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func UpdateUserRool(id uint, rool string) {
	var user models.User
	log.Println(id)

	db := handlers.Database()
	db.Find(&user, id)

	user.Rool = rool

	log.Println(rool)

	db.Save(&user)
	db.Close()
}

func UpdateUserName(id uint, last string, first string) {
	var user models.User

	db := handlers.Database()
	db.First(&user, id)

	user.FirstName = first
	user.LastName = last

	db.Save(&user)
	db.Close()
}

func UpdateUserPassword(id uint, password string, newPassword string) {
	var user models.User

	db := handlers.Database()
	db.First(&user, id)

	if user.Password != password {
		return
	}

	user.Password = newPassword

	db.Save(&user)
	db.Close()
}
