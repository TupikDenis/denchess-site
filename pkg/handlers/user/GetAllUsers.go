package user

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetAllUsers() []models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	for _, item := range users {
		_users = append(
			_users, models.TransformedUser{
				Id:        item.ID,
				Email:     item.Email,
				Password:  item.Password,
				LastName:  item.LastName,
				FirstName: item.FirstName,
				Rool:      item.Rool,
			})
	}

	return _users
}
