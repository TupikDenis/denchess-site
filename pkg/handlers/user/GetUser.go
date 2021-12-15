package user

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetUserByEmail(email string) models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser
	var user models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	db.Close()

	for _, item := range users {
		if item.Email == email {
			_users = append(
				_users, models.TransformedUser{
					Id:        item.ID,
					Email:     item.Email,
					Password:  item.Password,
					FirstName: item.FirstName,
					LastName:  item.LastName,
					Rool:      item.Rool,
				})
		}
	}

	if len(_users) == 1 {
		user = _users[0]
	}

	return user
}

func GetUserById(id uint) models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser
	var user models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	db.Close()

	for _, item := range users {
		if item.ID == id {
			_users = append(
				_users, models.TransformedUser{
					Id:        item.ID,
					Email:     item.Email,
					Password:  item.Password,
					FirstName: item.FirstName,
					LastName:  item.LastName,
					Rool:      item.Rool,
				})
		}
	}

	if len(_users) == 1 {
		user = _users[0]
	}

	return user
}

func Login(user_password string, user_email string) models.TransformedUser {
	var users []models.User
	var _users []models.TransformedUser
	var user models.TransformedUser

	db := handlers.Database()
	db.Find(&users)
	db.Close()

	for _, item := range users {
		if item.Email == user_email && item.Password == user_password {
			_users = append(
				_users, models.TransformedUser{
					Id:        item.ID,
					Email:     item.Email,
					Password:  item.Password,
					FirstName: item.FirstName,
					LastName:  item.LastName,
					Rool:      item.Rool,
				})
		}
	}

	if len(_users) == 1 {
		user = _users[0]
	}

	return user
}
