package redis

import (
	"context"
	"fmt"

	"base-gin-golang/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	*mongo.Client
}

func ConnectMongo(cfg *config.Environment) (*Database, error) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", cfg.MongoURI)),
	)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &Database{client}, nil
}
