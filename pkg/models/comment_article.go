package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CommentsArticle struct {
	gorm.Model
	IdArticle uint   `json:"id_news"`
	IdUser    uint   `json:"id_user"`
	UserName  string `json:"user_name"`
	Text      string `json:"text"`
}

type TransformedCommentsArticle struct {
	Id        uint      `json:"id"`
	IdArticle uint      `json:"id_news"`
	IdUser    uint      `json:"id_user"`
	UserName  string    `json:"user_name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
