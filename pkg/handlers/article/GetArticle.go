package articles

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetArticle(id uint) models.TransformedArticle {
	var articles []models.Article
	var _articles []models.TransformedArticle

	var readyArticles models.TransformedArticle

	db := handlers.Database()
	db.Find(&articles)
	//transforms the todos for building a good response
	for _, item := range articles {
		if item.ID == id {
			_articles = append(
				_articles, models.TransformedArticle{
					Id:          item.ID,
					IdUser:      item.IdUser,
					Title:       item.Title,
					Text:        item.Text,
					TextPreview: item.TextPreview,
					CreatedAt:   item.CreatedAt,
				})
			break
		}
	}

	readyArticles = _articles[len(_articles)-1]

	return readyArticles
}
