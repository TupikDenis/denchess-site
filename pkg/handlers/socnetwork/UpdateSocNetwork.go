package socnetwork

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func UpdateSocNetwork(id uint, instagram string, vk string) {
	var socnetwork models.SocNetwork

	db := handlers.Database()
	db.Where("").Find(&socnetwork)

	socnetwork.VK = vk
	socnetwork.Instagram = instagram

	db.Save(&socnetwork)

	db.Close()
}
