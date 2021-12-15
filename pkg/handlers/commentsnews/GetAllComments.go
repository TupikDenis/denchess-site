package commentsnews

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetAllComments(id uint) []models.TransformedCommentsNews {
	var comments []models.CommentsNews
	var _comments []models.TransformedCommentsNews

	db := handlers.Database()
	db.Find(&comments)
	for _, item := range comments {
		if item.IdNews == id {
			_comments = append(
				_comments, models.TransformedCommentsNews{
					Id:        item.ID,
					IdNews:    item.IdNews,
					IdUser:    item.IdUser,
					UserName:  item.UserName,
					Text:      item.Text,
					CreatedAt: item.CreatedAt,
				})
		}
	}

	return _comments
}
