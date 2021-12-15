package user

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func DeleteUser(user *models.User, id uint) {
	db := handlers.Database()
	db.First(&user, id)

	if id <= 0 {
		//редирект на ошибку
		return
	}

	db.Delete(&user)

}
