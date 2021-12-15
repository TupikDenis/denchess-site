package socnetwork

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetAllUserSocNetworkWithoutId() []models.TransformedSocNetwork {
	var socnetworks []models.SocNetwork
	var _socnetworks []models.TransformedSocNetwork

	db := handlers.Database()
	db.Find(&socnetworks)
	//transforms the todos for building a good response
	for _, item := range socnetworks {
		_socnetworks = append(
			_socnetworks, models.TransformedSocNetwork{
				Id:        item.ID,
				IdUser:    item.IdUser,
				Instagram: item.Instagram,
				VK:        item.VK,
			})
	}

	return _socnetworks
}

func GetAllUserSocNetwork(id uint) []string {
	var socnetwork []models.SocNetwork
	var _socnetwork []models.TransformedSocNetwork

	var socnetworks []string

	db := handlers.Database()
	db.Find(&socnetwork)
	//transforms the todos for building a good response
	for _, item := range socnetwork {
		if item.IdUser == id {
			_socnetwork = append(
				_socnetwork, models.TransformedSocNetwork{
					Id:        item.ID,
					IdUser:    item.IdUser,
					Instagram: item.Instagram,
					VK:        item.VK,
				})
		}
	}

	socnetworks = append(socnetworks, _socnetwork[len(_socnetwork)-1].Instagram)
	socnetworks = append(socnetworks, _socnetwork[len(_socnetwork)-1].VK)

	return socnetworks
}
