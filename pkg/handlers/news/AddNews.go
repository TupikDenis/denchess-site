package news

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddNews(idUser uint, title string, text string) {
	news := models.News{
		IdUser: idUser,
		Title:  title,
		Text:   text,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.News{})
	db.Create(&news)
	db.Close()
}
