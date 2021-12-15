package user

import (
	"net/http"

	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/socnetwork"

	"github.com/TupikDenis/denchess-site.git/pkg/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	user := models.User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Rool:      "user",
	}

	db := handlers.Database()
	db.AutoMigrate(&models.User{})
	db.Create(&user)

	socnetwork.AddSocNetwork(user.ID)

	db.Close()

	c.Redirect(http.StatusFound, "/signin")
}
