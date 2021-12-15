package articles

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func DeleteArticle(article *models.Article, id uint) {
	db := handlers.Database()
	db.First(&article, id)
	db.Delete(&article)
	db.Close()
}
