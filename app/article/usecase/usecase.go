package usecase

import (
	"github.com/apostrophedottilde/blog-article-api/app/article"
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
	"strconv"
	"time"
)

type usecase struct {
	repository article.Repository
}

func (usecase *usecase) Create(obj model.ArticleModel) (string, error) {
	nowInMillisString := strconv.Itoa(int(time.Now().UnixNano()))
	obj.Created = nowInMillisString
	obj.Updated = nowInMillisString
	savedEntity, err := usecase.repository.Create(obj)
	if err != nil {
		return savedEntity, err
	}
	return savedEntity, nil
}

func (usecase *usecase) FindOne(id string) (model.ArticleModel, error) {
	res, err := usecase.repository.FindOne(id)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *usecase) FindAll() ([]model.ArticleModel, error) {
	articles, err := usecase.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (usecase *usecase) Update(id string, model model.ArticleModel) error {
	nowInMillisString := strconv.Itoa(int(time.Now().UnixNano()))
	model.Updated = nowInMillisString
	err := usecase.repository.Update(id, model)
	if err != nil {
		return err
	}
	return nil
}

func New(repo article.Repository) *usecase {
	return &usecase{
		repository: repo,
	}
}
