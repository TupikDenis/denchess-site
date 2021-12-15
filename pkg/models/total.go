package models

import "github.com/jinzhu/gorm"

type Total struct {
	gorm.Model
	Rank              int    `json:"rank"`
	Year              uint   `json:"year"`
	Username          string `json:"username"`
	TotalFirstPlace   int    `json:"total_first_place"`
	TotalSecondtPlace int    `json:"total_second_place"`
	TotalThirdPlace   int    `json:"total_third_place"`
	TotalPoints       int    `json:"total_points"`
	TotalScore        int    `json:"total_score"`
}

type TransformedTotal struct {
	Id                uint   `json:"id"`
	Rank              int    `json:"rank"`
	Year              uint   `json:"year"`
	Username          string `json:"username"`
	TotalFirstPlace   int    `json:"total_first_place"`
	TotalSecondtPlace int    `json:"total_second_place"`
	TotalThirdPlace   int    `json:"total_third_place"`
	TotalPoints       int    `json:"total_points"`
	TotalScore        int    `json:"total_score"`
}
