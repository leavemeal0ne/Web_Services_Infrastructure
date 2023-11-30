package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	User       string
	Password   string
	DB         string
	Collection string
}

func Mongo(config MongoConfig) (*mongo.Collection, error) {
	url := fmt.Sprintf("mongodb://%s:%s@mongodb", config.User, config.Password)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	return client.Database(config.DB).Collection(config.Collection), nil
}
