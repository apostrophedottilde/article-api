package model

import (
	"encoding/xml"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestArticleModel_Response(t *testing.T) {
	type fields struct {
		ID      primitive.ObjectID
		Content string
		Created string
		Updated string
	}

	articleResponse := ArticleResponse{
		ID:      "5cd82bcb770963bafe888925",
		Content: "XML CONTENT",
		Created: "12345678123456789",
		Updated: "09876543210987654",
	}

	idObject, _ := primitive.ObjectIDFromHex("5cd82bcb770963bafe888925")

	tests := []struct {
		name   string
		fields fields
		want   ArticleResponse
	}{
		{
			"Should correctly map an ArticleModel to an ArticleResponse",
			fields{
				ID:      idObject,
				Content: "XML CONTENT",
				Created: "12345678123456789",
				Updated: "09876543210987654",
			},
			articleResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article := &ArticleModel{
				ID:      tt.fields.ID,
				Content: tt.fields.Content,
				Created: tt.fields.Created,
				Updated: tt.fields.Updated,
			}
			if got := article.Response(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleModel.Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleCollection_Response(t *testing.T) {
	type fields struct {
		XMLName  xml.Name
		Articles []ArticleModel
	}

	idObject, _ := primitive.ObjectIDFromHex("5cd82bcb770963bafe888221")

	var articleModels1 = ArticleModel{
		ID:      idObject,
		Content: "XML CONTENT 1",
		Created: "12345678123456789",
		Updated: "09876543210987654",
	}

	var articleModels2 = ArticleModel{
		ID:      idObject,
		Content: "XML CONTENT 2",
		Created: "12345678123456789",
		Updated: "09876543210987654",
	}

	var articleModels3 = ArticleModel{
		ID:      idObject,
		Content: "XML CONTENT 3",
		Created: "12345678123456789",
		Updated: "09876543210987654",
	}

	articleModels := []ArticleModel{articleModels1, articleModels2, articleModels3}

	articleCollection := ArticleCollection{
		XMLName:  xml.Name{},
		Articles: articleModels,
	}

	articleResponse1 := articleModels1.Response()
	articleResponse2 := articleModels2.Response()
	articleResponse3 := articleModels3.Response()

	articleResponses := []ArticleResponse{articleResponse1, articleResponse2, articleResponse3}

	articleCollectionResponse := ArticleCollectionResponse{
		XMLName:  articleCollection.XMLName,
		Articles: articleResponses,
	}

	tests := []struct {
		name   string
		fields fields
		want   ArticleCollectionResponse
	}{
		{
			"Should correctly map an ArticleCollection to an ArticleResponseCollection",
			fields{
				articleCollection.XMLName,
				articleModels,
			},
			articleCollectionResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := &ArticleCollection{
				XMLName:  tt.fields.XMLName,
				Articles: tt.fields.Articles,
			}
			if got := collection.Response(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleCollection.Response() = %v, want %v", got, tt.want)
			}
		})
	}
}
