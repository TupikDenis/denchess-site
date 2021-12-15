package commentsarticle

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func AddComments(idArticle uint, idUser uint, userName string, text string) {
	comment := models.CommentsArticle{
		IdArticle: idArticle,
		IdUser:    idUser,
		UserName:  userName,
		Text:      text,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.CommentsArticle{})
	db.Create(&comment)
	db.Close()
}
