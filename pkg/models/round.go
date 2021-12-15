package models

import "github.com/jinzhu/gorm"

type Tournament struct {
	gorm.Model
	Year  uint   `json:"id"`
	Round string `json:"round"`
	JSON  string `json:"json"`
}

type TransformedTournament struct {
	Id    uint   `json:"id"`
	Year  uint   `json:"year"`
	Round string `json:"round"`
	JSON  string `json:"json"`
}

type FullTournament struct {
	Rank       int    `json:"rank"`
	Score      int    `json:"score"`
	Rating     int    `json:"rating"`
	Username   string `json:"username"`
	Title      string `json:"title"`
	Perfomance int    `json:"perfomance"`
}
