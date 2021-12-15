package news

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetLast10News() []models.TransformedNews {
	var news []models.News
	var _news []models.TransformedNews

	db := handlers.Database()
	db.Order("id desc").Find(&news)
	k := 0
	for _, item := range news {
		if k < 3 {
			_news = append(
				_news, models.TransformedNews{
					Id:        item.ID,
					IdUser:    item.IdUser,
					Title:     item.Title,
					Text:      item.Text,
					CreatedAt: item.CreatedAt,
				})
			k++
		} else {
			break
		}
	}

	return _news
}

func GetAllNews() []models.TransformedNews {
	var news []models.News
	var _news []models.TransformedNews

	db := handlers.Database()
	db.Order("id desc").Find(&news)
	for _, item := range news {
		_news = append(
			_news, models.TransformedNews{
				Id:        item.ID,
				IdUser:    item.IdUser,
				Title:     item.Title,
				Text:      item.Text,
				CreatedAt: item.CreatedAt,
			})
	}

	return _news
}
