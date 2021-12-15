package articles

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func UpdateArticle(id uint, title string, preview string, text string) {
	var article models.Article

	db := handlers.Database()
	db.First(&article, id)

	article.Title = title
	article.TextPreview = preview
	article.Text = text

	db.Save(&article)

	db.Close()
}
