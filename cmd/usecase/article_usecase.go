package usecase

import (
	"github.com/4406arthur/clean-micro-article-service/pkg/entity"
)

//Logger interface ...
type Logger interface {
	Log(message string) error
}

//ArticleInteractor ...
type ArticleInteractor struct {
	ArticleRepository entity.ArticleRepository
	Logger            Logger
}

func (interactor *ArticleInteractor) Get(title string) (entity.Article, error) {

}

func (interactor *ArticleInteractor) List() ([]*entity.Article, error) {

}

func (interactor *ArticleInteractor) Create(article entity.Article) error {

}
