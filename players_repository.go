package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = os.Getenv("MONGO_CONNECTION_STRING")
var database = os.Getenv("MONGO_DATABASE")
var collection = os.Getenv("MONGO_COLLECTION")

func updatePlayers(players []PlayerData) error {
	context := context.Background()

	collection, err := getMongoCollection(context)
	if err != nil {
		return err
	}

	models := make([]mongo.WriteModel, 0)
	for _, player := range players {
		models = append(models, player.toUpsertModel())
	}

	options := options.BulkWrite().SetOrdered(false)

	_, err = collection.BulkWrite(context, models, options)

	return err
}

func getMongoCollection(c context.Context) (mongo.Collection, error) {
	options := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(c, options)
	if err != nil {
		return mongo.Collection{}, err
	}

	return *client.Database(database).Collection(collection), nil
}

func (p PlayerData) toUpsertModel() mongo.WriteModel {
	filter := bson.D{{Key: "_id", Value: p.Global.Uid}}
	replacement := bson.D{
		{Key: "name", Value: p.Global.Name},
		{Key: "platform", Value: p.Global.Platform},
		{Key: "level", Value: p.Global.Level},
	}

	model := mongo.NewReplaceOneModel()
	model.SetFilter(filter)
	model.SetReplacement(replacement)
	model.SetUpsert(true)

	return model
}
