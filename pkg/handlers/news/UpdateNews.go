package news

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func UpdateNews(id uint, title string, text string) {
	var news models.News

	db := handlers.Database()
	db.First(&news, id)

	news.Title = title
	news.Text = text

	db.Save(&news)

	db.Close()
}
