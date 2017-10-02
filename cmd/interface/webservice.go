package interfaces

import (
	"net/http"

	"github.com/4406arthur/clean-micro-article-service/pkg/entity"
)

type ArticleInteractor interface {
	Get(title string) (*entity.Article, error)
	List() ([]*entity.Article, error)
	Create(article entity.Article) error
}

type WebserviceHandler struct {
	ArticleInteractor ArticleInteractor
}

func (handler WebserviceHandler) ApiList(res http.ResponseWriter, req *http.Request) {
	articles, _ := handler.ArticleInteractor.List()
	for _, article := range articles {

	}
}

func (handler WebserviceHandler) ApiGet(res http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	article, _ := handler.ArticleInteractor.Get(title)
	return article
}

func (handler WebserviceHandler) ApiCreate(res http.ResponseWriter, req *http.Request) {

	handler.ArticleInteractor.Create()
	for _, article := range articles {

	}
}
