package commentsnews

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddComments(idNews uint, idUser uint, userName string, text string) {
	comment := models.CommentsNews{
		IdNews:   idNews,
		IdUser:   idUser,
		UserName: userName,
		Text:     text,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.CommentsNews{})
	db.Create(&comment)
	db.Close()
}
