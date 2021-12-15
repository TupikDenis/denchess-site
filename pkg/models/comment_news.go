package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CommentsNews struct {
	gorm.Model
	IdNews   uint   `json:"id_news"`
	IdUser   uint   `json:"id_user"`
	UserName string `json:"user_name"`
	Text     string `json:"text"`
}

type TransformedCommentsNews struct {
	Id        uint      `json:"id"`
	IdNews    uint      `json:"id_news"`
	IdUser    uint      `json:"id_user"`
	UserName  string    `json:"user_name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
