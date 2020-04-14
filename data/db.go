package data

import (
	"context"

	"github.com/BenLubar/memoize"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDatase(uri string, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

var MemoConnectToDatase = memoize.Memoize(connectToDatase).(func(uri string, dbName string) (*mongo.Database, error))
