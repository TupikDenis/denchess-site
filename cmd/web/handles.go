package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/TupikDenis/denchess-site.git/pkg/files/myjsons"
	articles "github.com/TupikDenis/denchess-site.git/pkg/handlers/article"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/commentsarticle"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/commentsnews"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/news"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/socnetwork"
	tournaments "github.com/TupikDenis/denchess-site.git/pkg/handlers/tournament"
	"github.com/TupikDenis/denchess-site.git/pkg/handlers/user"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func userSession(c *gin.Context) (uint, string) {
	session := sessions.Default(c)
	var id uint
	var fullName string

	i := session.Get("id")

	if i == nil {
		id = 0
	} else {
		id = session.Get("id").(uint)
	}

	session.Set("id", id)

	userInfo := user.GetUserById(id)
	fullName = userInfo.LastName + " " + userInfo.FirstName

	session.Set("full_name", fullName)
	session.Set("email", userInfo.Email)
	session.Set("rool", userInfo.Rool)
	session.Save()

	return id, fullName
}

func convertUInt(id string) uint {
	idu64, _ := strconv.ParseUint(id, 10, 32)
	idu := uint(idu64)
	return idu
}

func home(c *gin.Context) {
	id, fullName := userSession(c)
	last3Articles := articles.GetLast3lArticle()
	last10News := news.GetLast10News()

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":     "Главная страница",
		"id":        id,
		"full_name": fullName,
		"articles":  last3Articles,
		"news":      last10News,
	})
}

//новости

func getAllNews(c *gin.Context) {
	id, fullName := userSession(c)

	allNews := news.GetAllNews()

	//pages := c.Query("page", "1")

	c.HTML(http.StatusOK, "news.html", gin.H{
		"title":     "Новости",
		"id":        id,
		"full_name": fullName,
		"news":      allNews,
	})
}

func getNews(c *gin.Context) {
	id, fullName := userSession(c)
	idNews := c.Param("id")
	iduNews := convertUInt(idNews)

	myNews := news.GetNews(iduNews)
	title := myNews.Title
	text := myNews.Text
	idNewsUser := myNews.IdUser

	allComments := commentsnews.GetAllComments(iduNews)

	c.HTML(http.StatusOK, "news_item.html", gin.H{
		"title":      "Новость",
		"id":         id,
		"id_news":    iduNews,
		"id_user":    idNewsUser,
		"full_name":  fullName,
		"title_news": title,
		"news_text":  text,
		"comments":   allComments,
	})
}

func addNews(c *gin.Context) {
	id, fullName := userSession(c)
	c.HTML(http.StatusOK, "add_news.html", gin.H{
		"title":     "Создать новость",
		"id":        id,
		"full_name": fullName,
	})
}

func addNewsProcess(c *gin.Context) {
	id, _ := userSession(c)

	title := c.PostForm("title")
	text := c.PostForm("news-content")

	log.Println(text)

	news.AddNews(id, title, text)
	url := "/users/:" + strconv.FormatUint(uint64(id), 10)
	c.Redirect(http.StatusFound, url)
}

func updateNews(c *gin.Context) {
	id, fullName := userSession(c)
	idNews := c.Param("id")
	iduNews := convertUInt(idNews)

	myNews := news.GetNews(iduNews)
	title := myNews.Title
	text := myNews.Text

	c.HTML(http.StatusOK, "edit_news.html", gin.H{
		"title":      "Обновить новость",
		"id":         id,
		"id_news":    iduNews,
		"full_name":  fullName,
		"title_news": title,
		"text":       text,
	})
}

func updateNewsProcess(c *gin.Context) {
	title := c.PostForm("title")
	text := c.PostForm("news-content")
	readyText := strings.TrimSpace(text)

	idNews := c.Param("id")
	iduNews := convertUInt(idNews)

	log.Println(iduNews)
	log.Println(title)
	log.Println(readyText)

	news.UpdateNews(iduNews, title, text)
}

func deleteNews(c *gin.Context) {
	var _news models.News
	id := c.Param("id")
	idu := convertUInt(id)

	news.DeleteNews(&_news, idu)
}

//статьи

func getAllArticles(c *gin.Context) {
	id, fullName := userSession(c)

	allArticles := articles.GetAllArticle()
	c.HTML(http.StatusOK, "articles.html", gin.H{
		"title":     "Статьи",
		"id":        id,
		"full_name": fullName,
		"articles":  allArticles,
	})
}

func getArticle(c *gin.Context) {
	id, fullName := userSession(c)
	idArticle := c.Param("id")
	iduArticle := convertUInt(idArticle)

	myArticle := articles.GetArticle(iduArticle)
	title := myArticle.Title
	preview := myArticle.TextPreview
	text := myArticle.Text
	createdAt := myArticle.CreatedAt

	idArticleUser := myArticle.IdUser
	author := user.GetUserById(idArticleUser)
	authorName := author.LastName + " " + author.FirstName

	allComments := commentsarticle.GetAllComments(iduArticle)

	c.HTML(http.StatusOK, "article.html", gin.H{
		"title":         "Новость",
		"id":            id,
		"id_article":    iduArticle,
		"id_user":       idArticleUser,
		"author":        authorName,
		"full_name":     fullName,
		"title_article": title,
		"preview":       preview,
		"article_text":  text,
		"created":       createdAt,
		"comments":      allComments,
	})
}

func addArticle(c *gin.Context) {
	id, fullName := userSession(c)
	c.HTML(http.StatusOK, "add_article.html", gin.H{
		"title":     "Создать статью",
		"id":        id,
		"full_name": fullName,
	})
}

func addArticleProcess(c *gin.Context) {
	id, _ := userSession(c)

	title := c.PostForm("title")
	preview := c.PostForm("article-preview")
	text := c.PostForm("article-content")

	articles.AddArticle(id, title, preview, text)
	url := "/users/:" + strconv.FormatUint(uint64(id), 10)
	c.Redirect(http.StatusFound, url)
}

func updateArticle(c *gin.Context) {
	id, fullName := userSession(c)
	idArticle := c.Param("id")
	iduArticle := convertUInt(idArticle)

	myArticle := articles.GetArticle(iduArticle)
	title := myArticle.Title
	preview := myArticle.TextPreview
	readyPreview := strings.TrimSpace(preview)
	text := myArticle.Text

	c.HTML(http.StatusOK, "edit_article.html", gin.H{
		"title":         "Обновить статью",
		"id":            id,
		"id_article":    iduArticle,
		"full_name":     fullName,
		"title_article": title,
		"preview":       readyPreview,
		"text":          text,
	})
}

func updateArticleProcess(c *gin.Context) {
	title := c.PostForm("title")
	preview := c.PostForm("article-preview")
	text := c.PostForm("article-content")

	idArticle := c.Param("id")
	iduArticle := convertUInt(idArticle)

	articles.UpdateArticle(iduArticle, title, preview, text)
}

func deleteArticle(c *gin.Context) {
	var _article models.Article
	id := c.Param("id")
	idu := convertUInt(id)
	articles.DeleteArticle(&_article, idu)
}

//кубок канала

func denchessCup(c *gin.Context) {
	id, fullName := userSession(c)
	session := sessions.Default(c)
	rool := session.Get("rool")
	c.HTML(http.StatusOK, "denchess_cup.html", gin.H{
		"title":     "Кубок канала DenChess",
		"id":        id,
		"full_name": fullName,
		"rool":      rool,
	})
}

func tournament(c *gin.Context) {
	id, fullName := userSession(c)
	year := c.Param("year")

	uyear := convertUInt(year)

	players := tournaments.GetPlayersFromTournament(uyear)

	for i := 0; i < len(players); i++ {
		players[i].Rank = i + 1
	}

	c.HTML(http.StatusOK, "tournaments.html", gin.H{
		"title":     "Кубок канала DenChess " + year + ". Все таблицы",
		"id":        id,
		"full_name": fullName,
		"players":   players,
	})
}

func addTournamentProcess(c *gin.Context) {
	name := c.PostForm("name")
	file := c.PostForm("number")
	year := c.PostForm("year")

	uyear := convertUInt(year)

	url := myjsons.DownloadJSON(file)

	tournaments.AddTournament(url, name, uyear)

	yearTournament := tournaments.GetAllTournaments(uyear)
	var tournament []models.FullTournament
	var tournamentName []string
	for i := 0; i < len(yearTournament); i++ {
		tournamentName = append(tournamentName, yearTournament[i].Round)
		err := json.Unmarshal([]byte(yearTournament[i].JSON), &tournament)
		for k := 0; k < len(tournament); k++ {
			first := 0
			second := 0
			third := 0
			score := 0
			switch k {
			case 0:
				first = 1
				score = 10
			case 1:
				second = 1
				score = 6
			case 2:
				third = 1
				score = 4
			case 3:
				score = 3
			case 4:
				score = 2
			case 5:
				score = 1
			}

			tournaments.AddTotal(yearTournament[i].Year, tournament[k].Username, first, second, third, tournament[k].Score, score)
			first = 0
			second = 0
			third = 0
		}

		if err != nil {
			log.Println(err)
		}
	}

	c.Redirect(http.StatusFound, "/tournaments/"+year)
}

//
func getAllArticlesEditors(c *gin.Context) {
	id, fullName := userSession(c)
	c.HTML(http.StatusOK, "articles_editors.html", gin.H{
		"title":     "Авторы статей",
		"id":        id,
		"full_name": fullName,
	})
}

func getlArticlesEditors(c *gin.Context) {
	id, fullName := userSession(c)
	c.HTML(http.StatusOK, "articles_editors.html", gin.H{
		"title":     "Авторы статей",
		"id":        id,
		"full_name": fullName,
	})
}

//регистрация, авторизация и аунтетификация
func signin(c *gin.Context) {
	id, fullName := userSession(c)

	c.HTML(http.StatusOK, "signin.html", gin.H{
		"title":     "Авторизация",
		"id":        id,
		"full_name": fullName,
	})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func GetUser(c *gin.Context) {
	user_email := c.PostForm("email")
	user_password := c.PostForm("password")

	user := user.Login(user_password, user_email)

	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("full_name", user.FirstName+" "+user.LastName)
	session.Set("rool", user.Rool)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func GetAllUsers(c *gin.Context) {
	id, fullName := userSession(c)

	users := user.GetAllUsers()

	c.HTML(http.StatusOK, "all_users.html", gin.H{
		"title":     "Список пользователей",
		"id":        id,
		"full_name": fullName,
		"users":     users,
	})

}

func UpdateUser(c *gin.Context) {
	id, fullName := userSession(c)

	user := user.GetUserById(id)
	firstName := user.FirstName
	lastName := user.LastName

	socnetwork := socnetwork.GetAllUserSocNetwork(id)

	instagram := socnetwork[0]
	vk := socnetwork[1]

	c.HTML(http.StatusOK, "user_update.html", gin.H{
		"title":      "Редактирование профиля",
		"id":         id,
		"first_name": firstName,
		"last_name":  lastName,
		"full_name":  fullName,
		"instagram":  instagram,
		"vk":         vk,
	})

}

func UpdateName(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	lastName := c.PostForm("last_name")
	firstName := c.PostForm("first_name")

	user.UpdateUserName(idu, lastName, firstName)
}

func UpdatePassword(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	password := c.PostForm("password")
	newPassword := c.PostForm("new_password")

	user.UpdateUserPassword(idu, password, newPassword)

}

func UpdateSocNetwork(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	vk := c.PostForm("vk")
	insta := c.PostForm("insta")

	soc := socnetwork.GetAllUserSocNetworkWithoutId()

	for i := 0; i < len(soc); i++ {
		if soc[i].IdUser == idu {
			idu = soc[i].Id
			break
		}
	}

	if vk == "" {
		vk = "none"
	}

	if insta == "" {
		insta = "none"
	}

	socnetwork.UpdateSocNetwork(idu, insta, vk)
}

func UpdateRool(c *gin.Context) {
	id := c.Param("id")
	idu := convertUInt(id)

	rool := c.PostForm("rool")
	log.Println(idu)

	user.UpdateUserRool(idu, rool)
}

func DeleteUser(c *gin.Context) {
	var _user models.User
	id := c.Param("id")
	idu := convertUInt(id)

	user.DeleteUser(&_user, idu)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func registration(c *gin.Context) {
	id, fullName := userSession(c)

	c.HTML(http.StatusOK, "registration.html", gin.H{
		"title":     "Регистрация",
		"id":        id,
		"full_name": fullName,
	})
}

func profile(c *gin.Context) {
	session := sessions.Default(c)

	id, fullName := userSession(c)
	email := session.Get("email")
	rool := session.Get("rool")

	socnetwork := socnetwork.GetAllUserSocNetwork(id)

	instagram := socnetwork[0]
	vk := socnetwork[1]

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title":     fullName,
		"id":        id,
		"full_name": fullName,
		"email":     email,
		"instagram": instagram,
		"vk":        vk,
		"rool":      rool,
	})
}

func addCommentNews(c *gin.Context) {
	id, fullName := userSession(c)

	idNews := c.Param("id")
	iduNews := convertUInt(idNews)

	c.HTML(http.StatusOK, "add_news_comment.html", gin.H{
		"title":     "Оставить комментарий",
		"id":        id,
		"full_name": fullName,
		"id_news":   iduNews,
	})

}

func addCommentNewsProcess(c *gin.Context) {
	id, fullName := userSession(c)
	text := c.PostForm("comment-content")
	idNews := c.PostForm("id_news")
	iduNews := convertUInt(idNews)

	commentsnews.AddComments(iduNews, id, fullName, text)

	c.Redirect(http.StatusFound, "/news/"+strconv.FormatUint(uint64(iduNews), 10))
}

func addCommentArticle(c *gin.Context) {
	id, fullName := userSession(c)

	idArticle := c.Param("id")
	iduArticle := convertUInt(idArticle)

	c.HTML(http.StatusOK, "add_article_comment.html", gin.H{
		"title":      "Оставить комментарий",
		"id":         id,
		"full_name":  fullName,
		"id_article": iduArticle,
	})

}

func addCommentArticleProcess(c *gin.Context) {
	id, fullName := userSession(c)
	text := c.PostForm("comment-content")
	idArticle := c.PostForm("id_article")
	iduArticle := convertUInt(idArticle)

	commentsarticle.AddComments(iduArticle, id, fullName, text)

	c.Redirect(http.StatusFound, "/article/"+strconv.FormatUint(uint64(iduArticle), 10))
}

func changeStatus(c *gin.Context) {
	id, fullName := userSession(c)

	idUser := c.Param("id")
	iuUser := convertUInt(idUser)
	userActive := user.GetUserById(iuUser)

	c.HTML(http.StatusOK, "change_status.html", gin.H{
		"title":     "Обновсить статус",
		"id":        id,
		"rool":      userActive.Rool,
		"full_name": fullName,
		"iuUser":    iuUser,
	})
}
