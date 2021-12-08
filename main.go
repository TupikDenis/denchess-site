package denchess

// set GO111MODULE=on
// go run ./cmd/web/.

import (
	"log"
	"net/http"

	//"github.com/TupikDenis/denchess-site.git/pkg/handlers/user"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("./ui/html/*")
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Главная страница",
		})
	})

	/*router.HandleFunc("/users", user.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", user.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", user.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", user.DeleteUser).Methods(http.MethodDelete)*/
	router.Run(":8000")

	log.Println("API is running!")
}

func ServerStart() {

}

func main() {
	HandleFunc()
}
