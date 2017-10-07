package interfaces

import (
	"github.com/4406arthur/clean-micro-article-service/pkg/entity"
	"github.com/mitchellh/mapstructure"
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
	type articles []entity.Article

	docs, err := repo.dbHandler.GetAll()
	if err != nil {
		return nil, err
	}

	err := mapstructure.Decode(docs, &articles)
	if err != nil {
		panic(err)
	}

	return articles, nil
}

func (repo *DbArticleRepo) GetByTitle(title string) (*entity.Article, error) {
	type article entity.Article

	doc, err := repo.dbHandler.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	err := mapstructure.Decode(docs, &article)
	if err != nil {
		panic(err)
	}

	return article, nil
}

func (repo *DbArticleRepo) Store(article entity.Article) error {
	err := repo.dbHandler.Save(article)
	if err != nil {
		return err
	}
	return nil
}
