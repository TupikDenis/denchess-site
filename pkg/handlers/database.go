package handlers

import "github.com/jinzhu/gorm"

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3307)/denchess_site?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
