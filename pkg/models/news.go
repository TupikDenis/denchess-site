package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	IdUser uint   `json:"id_user"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type TransformedNews struct {
	Id        uint      `json:"id"`
	IdUser    uint      `json:"id_user"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
