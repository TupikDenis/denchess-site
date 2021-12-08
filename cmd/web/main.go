package main

// set GO111MODULE=on
// go run ./cmd/web/.

import (
	"log"

	//"github.com/TupikDenis/denchess-site.git/pkg/handlers/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MigrateDatabase() {
	//db := Database()
	//db.AutoMigrate(&Product{})
}

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3307)/sard")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func HandleFunc() {
	MigrateDatabase()

	router := gin.Default()
	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("ui/html/*")

	router.GET("/", home)
	router.GET("/nopage", pageNotFound)
	router.GET("/news", news)
	router.GET("/articles", articles)
	router.GET("/denchess_cup", denchessCup)
	router.GET("/articles_editors", articlesEditors)
	//router.GET("/article_editors/{id}", article_editors)
	router.GET("/login", login)
	router.GET("/registration", registration)
	router.GET("/profile", profile)

	/*router.GET("/users", user.GetAllUsers).Methods(http.MethodGet)
	router.GET("/users/{id}", user.GetUser).Methods(http.MethodGet)
	router.POST("/users", user.AddUser).Methods(http.MethodPost)
	router.PUT("/users/{id}", user.UpdateUser).Methods(http.MethodPut)
	router.DELETE("/users/{id}", user.DeleteUser).Methods(http.MethodDelete)*/
	router.Run(":8000")

	log.Println("API is running!")
}

func main() {
	HandleFunc()
}
