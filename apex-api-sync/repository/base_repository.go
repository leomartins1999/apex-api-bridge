package repository

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = os.Getenv("MONGO_CONNECTION_STRING")

var mongoClient = buildMongoClient(connectionString)

func getCollection(db string, col string) *mongo.Collection {
	return mongoClient.Database(db).Collection(col)
}

func getBulkWriteOptions() *options.BulkWriteOptions {
	return options.BulkWrite().SetOrdered(false)
}

func buildMongoClient(conString string) *mongo.Client {
	options := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		log.Fatalln("base_repository#buildMongoClient - Error establishing connection to db", err)
	}

	return client
}
