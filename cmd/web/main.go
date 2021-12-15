package main

// set GO111MODULE=on
// go run ./cmd/web/.

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// @title DenChess Site
// @version 1.0
// @description Api Server for my site

// @host localhost:8080
// @basePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func add(param1 int, param2 int) int {
	return param1 + param2
}

func HandleFunc() {
	store := cookie.NewStore([]byte("secret"))
	router := gin.Default()
	router.Use(sessions.Sessions("mysession", store))

	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("ui/html/*")

	router.GET("/", home)
	router.GET("/change_status/:id", changeStatus)

	router.GET("/add_news/:id", addNews)
	router.GET("/edit_news/:id", updateNews)

	newsHandl := router.Group("/news")
	{
		newsHandl.GET("", getAllNews)
		newsHandl.GET("/:id", getNews)
		newsHandl.POST("", addNewsProcess)
		newsHandl.PUT("/:id", updateNewsProcess)
		newsHandl.DELETE("/:id", deleteNews)
		newsHandl.GET("/comments/:id", addCommentNews)
		newsHandl.POST("/comments", addCommentNewsProcess)
	}

	//router.GET("/news/edit_comments/:id", updateNews)
	//router.PUT("/news/comments/:id", updateNewsProcess)
	//router.DELETE("/news/comments/:id", deleteNews)

	router.GET("/add_article/:id", addArticle)
	router.GET("/edit_article/:id", updateArticle)

	articleHandl := router.Group("/article")
	{
		articleHandl.GET("", getAllArticles)
		articleHandl.GET("/:id", getArticle)
		articleHandl.POST("", addArticleProcess)
		articleHandl.PUT("/:id", updateArticleProcess)
		articleHandl.DELETE("/:id", deleteArticle)
		articleHandl.GET("/comments/:id", addCommentArticle)
		articleHandl.POST("/comments", addCommentArticleProcess)
	}

	router.GET("/denchess_cup", denchessCup)
	router.GET("/tournaments/:year", tournament)
	router.POST("add_tournament", addTournamentProcess)
	//router.GET("/articles_editors", getAllArticlesEditors)
	//router.GET("/articles_editors/articles/:id", getArticlesEditors)
	//router.GET("/article_editors/{id}", article_editors)

	router.GET("/signin", signin)
	router.GET("/registration", registration)

	router.POST("/user", GetUser)
	router.GET("/update_user/:id", UpdateUser)

	users := router.Group("/users")
	{
		users.GET("/:id", profile)
		users.POST("", user.AddUser)
		users.GET("", GetAllUsers)
		users.PUT("/names/:id", UpdateName)
		users.PUT("/password/:id", UpdatePassword)
		users.PUT("/rool/:id", UpdateRool)
		//users.PUT("/:id/icon", UpdateUser)
		users.PUT("/socnetwork/:id", UpdateSocNetwork)
		users.DELETE("/:id", DeleteUser)
	}

	router.GET("/logout", logout)

	router.Run(":8080")
}

func main() {
	HandleFunc()
}
