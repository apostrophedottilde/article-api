package model

import (
	"encoding/xml"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content,omitempty"`
	Created string             `bson:"created,omitempty"`
	Updated string             `bson:"updated,omitempty"`
}

func (article *ArticleModel) Response() ArticleResponse {
	response := ArticleResponse{}
	response.ID = article.ID.Hex()
	response.Content = article.Content
	response.Created = article.Created
	response.Updated = article.Updated
	return response
}

func (collection *ArticleCollection) Response() ArticleCollectionResponse {
	response := ArticleCollectionResponse{}
	for _, article := range collection.Articles {
		response.Articles = append(response.Articles, article.Response())
	}
	response.XMLName = collection.XMLName
	return response
}

type ArticleCollection struct {
	XMLName  xml.Name
	Articles []ArticleModel
}

type ArticleCollectionResponse struct {
	XMLName  xml.Name          `xml:"articles-collection"`
	Articles []ArticleResponse `xml:"articles"`
}

type ArticleResponse struct {
	XMLName xml.Name `xml:"article"`
	ID      string   `xml:"id"`
	Content string   `xml:"content"`
	Created string   `xml:"created"`
	Updated string   `xml:"updated"`
}
