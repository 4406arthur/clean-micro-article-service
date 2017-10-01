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

//FetchAll ...
func (interactor *ArticleInteractor) FetchAll() (*entity.Article, error) {
	listArticle, err := interactor.ArticleRepository.FetchAll()
	if err != nil {
		interactor.Logger.Log(err.Error())
		return nil, err
	}

	return listArticle, nil
}
