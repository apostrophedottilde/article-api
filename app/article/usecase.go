package article

import (
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
)

type Usecase interface {
	Create(model model.ArticleModel) (string, error)
	FindOne(id string) (model.ArticleModel, error)
	FindAll() ([]model.ArticleModel, error)
	Update(id string, model model.ArticleModel) error
}
