package socnetwork

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddSocNetwork(idUser uint) {

	socNetwork := models.SocNetwork{
		IdUser:    idUser,
		Instagram: "none",
		VK:        "none",
	}

	db := handlers.Database()
	db.AutoMigrate(&models.SocNetwork{})
	db.Create(&socNetwork)
	db.Close()
}
