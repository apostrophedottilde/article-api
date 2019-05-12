package mongo

import (
	"context"
	"errors"
	"github.com/apostrophedottilde/blog-article-api/app/article/model"
	"github.com/apostrophedottilde/blog-article-api/app/article/repository/mongo/wrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

var (
	dbName             = "blog"
	documentCollection = "articles"
)

type ArticleRepository struct {
	client wrapper.MongoDAL
}

func (repo *ArticleRepository) Create(entity model.ArticleModel) (string, error) {
	nowInMillisString := strconv.Itoa(int(time.Now().UnixNano()))
	entity.Created = nowInMillisString
	entity.Updated = nowInMillisString
	res, err := repo.client.InsertOne(entity)

	if err != nil {
		return "", errors.New("ERROR PERSISTING NEW ARTICLE")
	}

	hexid := res.InsertedID.(primitive.ObjectID).Hex()
	docID, err := primitive.ObjectIDFromHex(hexid)

	return docID.Hex(), nil
}

func (repo *ArticleRepository) FindOne(id string) (*model.ArticleModel, error) {
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("ERROR PARSING ARTICLE ID - MALFORMED")
	}

	filter := bson.D{{"_id", docID}}

	var result model.ArticleModel
	err = repo.client.FindOne(filter).Decode(&result)

	if err != nil {
		return nil, errors.New("ERROR NOT FOUND ARTICLE")
	}

	return &result, nil
}

func (repo *ArticleRepository) FindAll() ([]model.ArticleModel, error) {
	cur, ctx, err := repo.client.Find(bson.M{})
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

func (repo *ArticleRepository) Update(id string, updated model.ArticleModel) error {
	nowInMillisString := strconv.Itoa(int(time.Now().UnixNano()))
	updated.Updated = nowInMillisString

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ERROR PARSING ARTICLE ID - MALFORMED")
	}
	filter := bson.D{{"_id", docID}}

	var result model.ArticleModel
	err = repo.client.FindOne(filter).Decode(&result)
	result.Content = updated.Content
	result.Updated = updated.Updated

	if err != nil {
		return errors.New("ERROR NOT FOUND ARTICLE - CAN'T UPDATE ANYTHING")
	}
	replace := repo.client.FindOneAndReplace(filter, result)
	if replace == nil {
		return errors.New("ERROR UPDATING ARTICLE")
	}
	return nil
}

func NewRepository(client wrapper.MongoDAL) *ArticleRepository {
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
		client: client,
	}
}
