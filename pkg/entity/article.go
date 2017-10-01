package entity

import (
	"time"
)

//ArticleRepository ...
//A repository is a concept from Domain-driven Design
//it abstracts away the fact that domain entities need to be saved to or loaded from some kind of persistence mechanism
type ArticleRepository interface {
	FetchAll() ([]*Article, error)
	FindByTitle(title string) (Article, error)
	Store(article Article) error
}

//Article init entity
type Article struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
