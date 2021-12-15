package news

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func DeleteNews(news *models.News, id uint) {
	db := handlers.Database()
	db.First(&news, id)
	db.Delete(&news)
	db.Close()
}
