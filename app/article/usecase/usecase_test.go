package usecase

import (
	"errors"
	"github.com/apostrophedottilde/blog-article-api/app/article"
	"github.com/apostrophedottilde/blog-article-api/app/article/mocks"
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func Test_usecase_Create(t *testing.T) {
	type fields struct {
		repository article.Repository
	}
	type args struct {
		obj model.ArticleModel
	}

	mockRepository := &mocks.Repository{}
	mockFailureRepository := &mocks.Repository{}

	proposedArticle := model.ArticleModel{
		Content: "<xml><content>PROPOSED ARTICLE XML BODY</Content></xml>",
	}

	mockRepository.On("Create", proposedArticle).Return("5cd6bb192c6b2e7e5126a3ef", nil)
	mockFailureRepository.On("Create", proposedArticle).Return(nil, errors.New("ERROR PERSISTING NEW ARTICLE"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"Should return the ID of the newly created Article as a string",
			fields{mockRepository},
			args{proposedArticle},
			"5cd6bb192c6b2e7e5126a3ef",
			false,
		},
		{
			"Should return an error message when the Repository call fails",
			fields{mockRepository},
			args{proposedArticle},
			"5cd6bb192c6b2e7e5126a3ef",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &usecase{
				repository: tt.fields.repository,
			}
			got, err := usecase.Create(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_FindOne(t *testing.T) {
	type fields struct {
		repository article.Repository
	}
	type args struct {
		id string
	}

	mockRepository := &mocks.Repository{}

	id, _ := primitive.ObjectIDFromHex("5cd82bca770963bafe88891f")

	fetchedArticle := model.ArticleModel{
		ID:      id,
		Content: "<xml><content>PROPOSED ARTICLE XML BODY</Content></xml>",
	}

	mockRepository.On("FindOne", "5cd82bca770963bafe88891f").Return(fetchedArticle, nil)
	mockRepository.On("FindOne", "INVALID_ID").Return(model.ArticleModel{}, errors.New("ERROR PARSING ARTICLE ID - MALFORMED"))
	mockRepository.On("FindOne", "5cd82bca770963bafe888921").Return(model.ArticleModel{}, errors.New("ERROR NOT FOUND ARTICLE"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.ArticleModel
		wantErr bool
	}{
		{
			"Should return an Article entity on successful Repository invocation",
			fields{mockRepository},
			args{"5cd82bca770963bafe88891f"},
			fetchedArticle,
			false,
		},
		{
			"Should return an error message when the ID is badly formed",
			fields{mockRepository},
			args{"INVALID_ID"},
			model.ArticleModel{},
			true,
		},
		{
			"Should return an error message when the Article is not found",
			fields{mockRepository},
			args{"5cd82bca770963bafe888921"},
			model.ArticleModel{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &usecase{
				repository: tt.fields.repository,
			}
			got, err := usecase.FindOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.FindOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_FindAll(t *testing.T) {
	type fields struct {
		repository article.Repository
	}

	mockRepository := &mocks.Repository{}
	mockFailureRepository := &mocks.Repository{}

	id, _ := primitive.ObjectIDFromHex("5cd82bca770963bafe88891f")

	fetchedArticles := []model.ArticleModel{{
		ID:      id,
		Content: "<xml><content>BODY 1</Content></xml>",
	}, {
		ID:      id,
		Content: "<xml><content>BODY 2</Content></xml>",
	}, {
		ID:      id,
		Content: "<xml><content>BODY 3</Content></xml>",
	},
	}

	mockRepository.On("FindAll").Return(fetchedArticles, nil)
	mockFailureRepository.On("FindAll").Return([]model.ArticleModel{}, errors.New("ERROR FETCHING COLLECTION OF ARTICLES"))

	tests := []struct {
		name    string
		fields  fields
		want    []model.ArticleModel
		wantErr bool
	}{
		{
			"Should return list of all Articles from Repository",
			fields{mockRepository},
			fetchedArticles,
			false,
		},
		{
			"On error making invocation on repository should return the error",
			fields{mockFailureRepository},
			[]model.ArticleModel{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &usecase{
				repository: tt.fields.repository,
			}
			got, err := usecase.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_Update(t *testing.T) {
	type fields struct {
		repository article.Repository
	}
	type args struct {
		id    string
		model model.ArticleModel
	}

	mockRepository := &mocks.Repository{}
	mockFailureRepository := &mocks.Repository{}

	id, _ := primitive.ObjectIDFromHex("5cd82bca770963bafe88891f")

	articleUpdate := model.ArticleModel{
		ID:      id,
		Content: "<xml><content>PROPOSED ARTICLE XML BODY</Content></xml>",
	}

	mockRepository.On("Update", "5cd82bca770963bafe88891f", articleUpdate).Return(nil)
	mockFailureRepository.On("Update", "INVALID_ID", articleUpdate).Return(errors.New("ERROR PARSING ARTICLE ID - MALFORMED"))
	mockFailureRepository.On("Update", "5cd82bca770963bafe88891f", articleUpdate).Return(errors.New("ERROR UPDATING ARTICLE"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Should not return an error when update is successful",
			fields{mockRepository},
			args{"5cd82bca770963bafe88891f", articleUpdate},
			false,
		},
		{
			"Should return an error when the ID if malformed",
			fields{mockFailureRepository},
			args{"INVALID_ID", articleUpdate},
			true,
		},
		{
			"Should return an error when there is a problem updating",
			fields{mockFailureRepository},
			args{"5cd82bca770963bafe88891f", articleUpdate},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &usecase{
				repository: tt.fields.repository,
			}
			if err := usecase.Update(tt.args.id, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("usecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		repo article.Repository
	}

	mockRepository := &mocks.Repository{}

	expectedUseCase := usecase{mockRepository}

	tests := []struct {
		name string
		args args
		want *usecase
	}{
		{
			"Should return an error when there is a problem updating",
			args{mockRepository},
			&expectedUseCase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
