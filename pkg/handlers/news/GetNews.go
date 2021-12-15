package news

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetNews(id uint) models.TransformedNews {
	var news []models.News
	var _news []models.TransformedNews

	var readyNews models.TransformedNews

	db := handlers.Database()
	db.Find(&news)
	//transforms the todos for building a good response
	for _, item := range news {
		if item.ID == id {
			_news = append(
				_news, models.TransformedNews{
					Id:        item.ID,
					IdUser:    item.IdUser,
					Title:     item.Title,
					Text:      item.Text,
					CreatedAt: item.CreatedAt,
				})
			break
		}
	}

	readyNews = _news[len(_news)-1]

	return readyNews
}
