package pkg

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Birthday  string `json:"birthday"`
}

type TransformUser struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Birthday  string `json:"birthday"`
}
