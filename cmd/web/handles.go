package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Главная страница",
	})
}

func pageNotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "nopage.html", gin.H{
		"title": "Страница не найдена",
	})
}

func news(c *gin.Context) {
	c.HTML(http.StatusOK, "news.html", gin.H{
		"title": "Новости",
	})
}

func articles(c *gin.Context) {
	c.HTML(http.StatusOK, "articles.html", gin.H{
		"title": "Статьи",
	})
}

func denchessCup(c *gin.Context) {
	c.HTML(http.StatusOK, "articles.html", gin.H{
		"title": "Кубок канала DenChess",
	})
}

func articlesEditors(c *gin.Context) {
	c.HTML(http.StatusOK, "articles_editors.html", gin.H{
		"title": "Авторы статей",
	})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Авторизация",
	})
}

func registration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", gin.H{
		"title": "Регистрация",
	})
}

func profile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title": "Профиль",
	})
}
