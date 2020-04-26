package src

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func getCollection() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"))
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil
	}
	collection := client.Database("db").Collection("coll")

	return collection
}