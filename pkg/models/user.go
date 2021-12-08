package models

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}
