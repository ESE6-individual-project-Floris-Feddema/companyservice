package contexts

import (
	"companyservice/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection = nil;
var client *mongo.Client = nil;

func GetCollection(ctx context.Context) *mongo.Collection {
	if collection != nil {
		return collection
	}

	connectionString := utils.EnvVar("DB_CONNECTION_STRING")
	databaseName := utils.EnvVar("COMPANY_DATABASE")
	collectionName := utils.EnvVar("COMPANY_COLLECTION")

	clientOptions := options.Client().ApplyURI(connectionString)
	clientConn, err := mongo.Connect(ctx, clientOptions)

	client = clientConn

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(databaseName).Collection(collectionName)

	return collection
}

func GetClient(ctx context.Context) *mongo.Client {
	if collection == nil {
		GetCollection(ctx)
	}
	return client
}