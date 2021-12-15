package commentsarticle

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetAllComments(id uint) []models.TransformedCommentsArticle {
	var comments []models.CommentsArticle
	var _comments []models.TransformedCommentsArticle

	db := handlers.Database()
	db.Find(&comments)
	for _, item := range comments {
		if item.IdArticle == id {
			_comments = append(
				_comments, models.TransformedCommentsArticle{
					Id:        item.ID,
					IdArticle: item.IdArticle,
					IdUser:    item.IdUser,
					UserName:  item.UserName,
					Text:      item.Text,
					CreatedAt: item.CreatedAt,
				})
		}
	}

	return _comments
}
