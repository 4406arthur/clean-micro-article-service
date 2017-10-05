package interfaces

import (
	"github.com/4406arthur/clean-micro-article-service/pkg/entity"
)

//TODO
type DbHandler interface {
	GetAll() ([]map[string]interface{}, error)
	GetByTitle(title string) (map[string]interface{}, error)
	Save(article map[string]interface{}) error
}

type DbRepo struct {
	dbHandler DbHandler
}

type DbArticleRepo DbRepo

func NewDbArticleRepo(dbHandler DbHandler) *DbArticleRepo {
	dbArticleRepo := new(DbArticleRepo)
	dbArticleRepo.dbHandler = dbHandler
	return dbArticleRepo
}

func (repo *DbArticleRepo) FetchAll() ([]*entity.Article, error) {
	articles, err := repo.dbHandler.GetAll()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (repo *DbArticleRepo) GetByTitle(title string) (*entity.Article, error) {
	article, err := repo.dbHandler.GetByTitle(title)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (repo *DbArticleRepo) Store(article entity.Article) error {
	err := repo.dbHandler.Save(article)
	if err != nil {
		return err
	}
	return nil
}
