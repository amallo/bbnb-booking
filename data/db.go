package data

import (
	"context"
	"time"

	"github.com/BenLubar/memoize"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDatase(uri string, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

var MemoConnectToDatase = memoize.Memoize(connectToDatase).(func(uri string, dbName string) (*mongo.Database, error))
