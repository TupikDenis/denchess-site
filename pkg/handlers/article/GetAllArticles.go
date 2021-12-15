package articles

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetLast3lArticle() []models.TransformedArticle {
	var articles []models.Article
	var _articles []models.TransformedArticle

	db := handlers.Database()
	db.Order("id desc").Find(&articles)
	k := 0
	for _, item := range articles {
		if k < 3 {
			_articles = append(
				_articles, models.TransformedArticle{
					Id:          item.ID,
					IdUser:      item.IdUser,
					Title:       item.Title,
					Text:        item.Text,
					TextPreview: item.TextPreview,
					CreatedAt:   item.CreatedAt,
				})
			k++
		} else {
			break
		}
	}

	return _articles
}

func GetAllArticle() []models.TransformedArticle {
	var articles []models.Article
	var _articles []models.TransformedArticle

	db := handlers.Database()
	db.Order("id desc").Find(&articles)
	for _, item := range articles {
		_articles = append(
			_articles, models.TransformedArticle{
				Id:          item.ID,
				IdUser:      item.IdUser,
				Title:       item.Title,
				Text:        item.Text,
				TextPreview: item.TextPreview,
				CreatedAt:   item.CreatedAt,
			})
	}

	return _articles
}
