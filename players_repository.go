package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

func fetchUIDs() ([]string, error) {
	context := context.Background()

	collection, err := getMongoCollection(context)
	if err != nil {
		return []string{}, err
	}

	results, err := collection.Distinct(context, "_id", bson.D{})
	if err != nil {
		return []string{}, err
	}

	uids := make([]string, 0)
	for _, uid := range results {
		uids = append(uids, fmt.Sprintf("%v", uid))
	}

	return uids, nil
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
	filter := bson.D{{Key: "_id", Value: fmt.Sprint(p.Global.Uid)}}
	replacement := bson.D{
		{Key: "name", Value: p.Global.Name},
		{Key: "platform", Value: p.Global.Platform},
		{Key: "level", Value: p.Global.Level},
		{Key: "rank", Value: p.Global.getRank()},
		{Key: "rankPoints", Value: p.Global.Rank.RankScore},
		{Key: "updatedAt", Value: time.Now()},
	}

	model := mongo.NewReplaceOneModel()
	model.SetFilter(filter)
	model.SetReplacement(replacement)

	return model
}
