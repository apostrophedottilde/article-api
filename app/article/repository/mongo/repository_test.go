package mongo

//
//import (
//	"context"
//	"github.com/apostrophedottilde/blog-article-api/app/article/mocks"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"testing"
//
//	"github.com/apostrophedottilde/blog-article-api/app/article/model"
//	"go.mongodb.org/mongo-driver/mongo"
//	"github.com/stretchr/testify/assert"
//)
//
//type mockcMongoClientConnection struct {
//
//}
//
//type mockMongoClient struct {
//	mongo.Client
//}
//
//func (mock *mockcMongoClientConnection) InsertOne(ctx context.Context, articleModel model.ArticleModel) (mongo.InsertOneResult, error) {
//	newArticleId, _ := primitive.ObjectIDFromHex("5cd6bb182c6b2e7e5126a3ed")
//	return mongo.InsertOneResult{InsertedID: newArticleId}, nil
//}
//
//func TestArticleRepository_Create(t *testing.T) {
//
//	mockMongoClient{}.Database().Collection().
//	var newArticle = model.ArticleModel{
//		Content: "<XML><CONTENT>Mock XML string content<CONTENT></XML>",
//	}
//
//	mocks.Repository{}
//
//	type fields struct {
//		client interface{}
//	}
//
//	type args struct {
//		entity model.ArticleModel
//	}
//
//	tests := []struct {
//		name    string
//		dbClient  mongo.Client
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{"Should return the id of the successfully persisted entity.",
//			mockMongoClient{}.Client,
//			args{newArticle},
//			"5cd6bb182c6b2e7e5126a3ed",
//			false,
//		},
//
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ps := &ArticleRepository{
//				client: tt.dbClient,
//			}
//			got, err := ps.Create(tt.args.entity)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ArticleRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("ArticleRepository.Create() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
////func TestArticleRepository_FindOne(t *testing.T) {
////	type fields struct {
////		client mongo.Client
////	}
////	tests := []struct {
////		name    string
////		fields  fields
////		want    model.ArticleModel
////		wantErr bool
////	}{
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			ps := &ArticleRepository{
////				client: tt.fields.client,
////			}
////			got, err := ps.FindOne(tt.args.id)
////			if (err != nil) != tt.wantErr {
////				t.Errorf("ArticleRepository.FindOne() error = %v, wantErr %v", err, tt.wantErr)
////				return
////			}
////			if !reflect.DeepEqual(got, tt.want) {
////				t.Errorf("ArticleRepository.FindOne() = %v, want %v", got, tt.want)
////			}
////		})
////	}
////}
////
////func TestArticleRepository_FindAll(t *testing.T) {
////	type fields struct {
////		client mongo.Client
////	}
////	tests := []struct {
////		name    string
////		fields  fields
////		want    []model.ArticleModel
////		wantErr bool
////	}{
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			ps := &ArticleRepository{
////				client: tt.fields.client,
////			}
////			got, err := ps.FindAll()
////			if (err != nil) != tt.wantErr {
////				t.Errorf("ArticleRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
////				return
////			}
////			if !reflect.DeepEqual(got, tt.want) {
////				t.Errorf("ArticleRepository.FindAll() = %v, want %v", got, tt.want)
////			}
////		})
////	}
////}
////
////func TestArticleRepository_Update(t *testing.T) {
////	type fields struct {
////		client mongo.Client
////	}
////	type args struct {
////		id      string
////		updated model.ArticleModel
////	}
////	tests := []struct {
////		name    string
////		fields  fields
////		args    args
////		wantErr bool
////	}{
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			ps := &ArticleRepository{
////				client: tt.fields.client,
////			}
////			if err := ps.Update(tt.args.id, tt.args.updated); (err != nil) != tt.wantErr {
////				t.Errorf("ArticleRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
////			}
////		})
////	}
////}
////
////func TestNewRepository(t *testing.T) {
////	tests := []struct {
////		name string
////		want *ArticleRepository
////	}{
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			if got := NewRepository(); !reflect.DeepEqual(got, tt.want) {
////				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
////			}
////		})
////	}
////}
