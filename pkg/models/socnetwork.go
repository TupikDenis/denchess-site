package models

import "github.com/jinzhu/gorm"

type SocNetwork struct {
	gorm.Model
	IdUser    uint   `json:"id_user"`
	Instagram string `json:"instagram"`
	VK        string `json:"vk"`
}

type TransformedSocNetwork struct {
	Id        uint   `json:"id"`
	IdUser    uint   `json:"id_user"`
	Instagram string `json:"instagram"`
	VK        string `json:"vk"`
}
