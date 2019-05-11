package mongo

import (
	"context"
	"errors"
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	dbName             = "blog"
	documentCollection = "articles"
)

type ArticleRepository struct {
	client mongo.Client
}

func (ps *ArticleRepository) Create(entity model.ArticleModel) (string, error) {
	collection := ps.client.Database(dbName).Collection(documentCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	res, err := collection.InsertOne(ctx, entity)
	defer cancel()

	if err != nil {
		return "", errors.New("ERROR PERSISTING NEW ARTICLE")
	}

	hexid := res.InsertedID.(primitive.ObjectID).Hex()
	docID, err := primitive.ObjectIDFromHex(hexid)

	return docID.Hex(), nil
}

func (ps *ArticleRepository) FindOne(id string) (model.ArticleModel, error) {
	collection := ps.client.Database(dbName).Collection(documentCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.ArticleModel{}, errors.New("ERROR PARSING ARTICLE ID - MALFORMED")
	}

	filter := bson.D{{"_id", docID}}

	var result model.ArticleModel
	err = collection.FindOne(ctx, filter).Decode(&result)
	defer cancel()

	if err != nil {
		return result, errors.New("ERROR NOT FOUND ARTICLE")
	}

	return result, nil
}

func (ps *ArticleRepository) FindAll() ([]model.ArticleModel, error) {
	collection := ps.client.Database(dbName).Collection(documentCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	defer cancel()

	if err != nil {
		return nil, errors.New("ERROR FETCHING COLLECTION OF ARTICLES")
	}
	var projects []model.ArticleModel

	for cur.Next(ctx) {
		var result model.ArticleModel
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, result)
	}
	return projects, nil
}

func (ps *ArticleRepository) Update(id string, updated model.ArticleModel) error {
	collection := ps.client.Database(dbName).Collection(documentCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ERROR PARSING ARTICLE ID - MALFORMED")
	}
	filter := bson.D{{"_id", docID}}

	var result model.ArticleModel
	err = collection.FindOne(ctx, filter).Decode(&result)
	result.Content = updated.Content
	result.Updated = updated.Updated

	if err != nil {
		return errors.New("ERROR NOT FOUND ARTICLE - CAN'T UPDATE ANYTHING")
	}
	replace := collection.FindOneAndReplace(ctx, filter, result)
	if replace == nil {
		return errors.New("ERROR UPDATING ARTICLE")
	}
	return nil
}

func NewRepository() *ArticleRepository {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		errors.New("ERROR CONNECTING TO DATABASE")
	}

	err = mClient.Ping(context.TODO(), nil)
	if err != nil {
		errors.New("ERROR PINGING DATABASE")
	}
	return &ArticleRepository{
		client: *mClient,
	}
}
