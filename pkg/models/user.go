package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Rool      string `json:"rool"`
}

type TransformedUser struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Rool      string `json:"rool"`
}
