package wrapper

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	dbName             = "blog"
	documentCollection = "articles"
)

type MongoDAL interface {
	InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, context.Context, error)
	FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
}

type MongoWrapper struct {
	client mongo.Client
}

func (wrapper *MongoWrapper) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := wrapper.client.Database(dbName).Collection(documentCollection)
	res, err := collection.InsertOne(ctx, document)
	return res, err
}

func (wrapper *MongoWrapper) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := wrapper.client.Database(dbName).Collection(documentCollection)
	res := collection.FindOne(ctx, filter)
	defer cancel()
	return res
}

func (wrapper *MongoWrapper) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, context.Context, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := wrapper.client.Database(dbName).Collection(documentCollection)
	cur, err := collection.Find(ctx, bson.M{})
	return cur, ctx, err
}

func (wrapper *MongoWrapper) FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	collection := wrapper.client.Database(dbName).Collection(documentCollection)
	defer cancel()
	replace := collection.FindOneAndReplace(ctx, filter, replacement)
	return replace
}

func NewMongoWrapper() *MongoWrapper {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))

	if err != nil {
		errors.New("ERROR CONNECTING TO DATABASE")
	}

	err = mClient.Ping(context.TODO(), nil)
	if err != nil {
		errors.New("ERROR PINGING DATABASE")
	}
	return &MongoWrapper{
		client: *mClient,
	}
}
