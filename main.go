package main

import (
	"net/http"

	"github.com/4406arthur/clean-micro-article-service/cmd/infrastruct"
	"github.com/4406arthur/clean-micro-article-service/cmd/interfaces"
	"github.com/4406arthur/clean-micro-article-service/cmd/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	dbHandler := infrastruct.NewMongoHandler(localhost)

	articleInteractor := new(usecase.ArticleInteractor)
	articleInteractor.ArticleRepository = interfaces.NewDbArticleRepo(dbHandler)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ArticleInteractor = webserviceHandler

	router := gin.Default()

	router.GET("/article/:title", func(c *gin.Context) {
		title := c.Param("title")
		article := articleInteractor.Get(title)

		c.String(http.StatusOK, article)
	})

	router.Run(":8080")
}
