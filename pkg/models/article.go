package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	IdUser      uint   `json:"id_user"`
	Title       string `json:"title"`
	TextPreview string `json:"text_preview"`
	Text        string `json:"text"`
}

type TransformedArticle struct {
	Id          uint      `json:"id"`
	IdUser      uint      `json:"id_user"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	TextPreview string    `json:"text_preview"`
	CreatedAt   time.Time `json:"created_at"`
}
