package articles

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddArticle(idUser uint, title string, preview string, text string) {
	article := models.Article{
		IdUser:      idUser,
		Title:       title,
		TextPreview: preview,
		Text:        text,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.Article{})
	db.Create(&article)
	db.Close()
}
