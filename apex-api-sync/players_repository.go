package main

import (
	"apex-api-sync/models"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = os.Getenv("MONGO_CONNECTION_STRING")

var playersDB = os.Getenv("MONGO_PLAYERS_DATABASE")
var playersCol = os.Getenv("MONGO_PLAYERS_COLLECTION")

var gamesDB = os.Getenv("MONGO_GAMES_DATABASE")
var gamesCol = os.Getenv("MONGO_GAMES_COLLECTION")

func updatePlayers(players []models.PlayerData) error {
	context := context.Background()

	collection, err := getMongoCollection(context, playersDB, playersCol)
	if err != nil {
		return err
	}

	models := make([]mongo.WriteModel, 0)
	for _, player := range players {
		models = append(models, playerDataToWriteModel(player))
	}

	options := options.BulkWrite().SetOrdered(false)

	_, err = collection.BulkWrite(context, models, options)

	return err
}

func fetchUIDs() ([]string, error) {
	context := context.Background()

	collection, err := getMongoCollection(context, playersDB, playersCol)
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

func updateGames(games []models.GameData) error {
	context := context.Background()

	collection, err := getMongoCollection(context, gamesDB, gamesCol)
	if err != nil {
		return err
	}

	models := make([]mongo.WriteModel, 0)
	for _, game := range games {
		models = append(models, gameDataToWriteModel(game))
	}

	options := options.BulkWrite().SetOrdered(false)

	_, err = collection.BulkWrite(context, models, options)

	return err
}

func getMongoCollection(c context.Context, databaseName string, collectionName string) (mongo.Collection, error) {
	options := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(c, options)
	if err != nil {
		return mongo.Collection{}, err
	}

	return *client.Database(databaseName).Collection(collectionName), nil
}

func playerDataToWriteModel(p models.PlayerData) mongo.WriteModel {
	filter := bson.D{{Key: "_id", Value: fmt.Sprint(p.Global.Uid)}}
	replacement := bson.D{
		{Key: "name", Value: p.Global.Name},
		{Key: "platform", Value: p.Global.Platform},
		{Key: "level", Value: p.Global.Level},
		{Key: "rank", Value: p.Global.GetRank()},
		{Key: "rankPoints", Value: p.Global.Rank.RankScore},
		{Key: "updatedAt", Value: time.Now()},
		{Key: "selectedLegend", Value: p.Realtime.SelectedLegend},
	}

	model := mongo.NewReplaceOneModel()
	model.SetFilter(filter)
	model.SetReplacement(replacement)

	return model
}

func gameDataToWriteModel(g models.GameData) mongo.WriteModel {
	filter := bson.D{{Key: "_id", Value: g.GetGameId()}}
	replacement := bson.D{
		{Key: "playerId", Value: g.PlayerUid},
		{Key: "playerName", Value: g.PlayerName},
		{Key: "gameMode", Value: g.GameMode},
		{Key: "legendPlayed", Value: g.LegendPlayed},
		{Key: "startedAt", Value: time.Unix(g.StartTimestamp, 0)},
		{Key: "endedAt", Value: time.Unix(g.EndTimestamp, 0)},
		{Key: "damageDone", Value: g.GetDamageDone()},
		{Key: "kills", Value: g.GetKills()},
		{Key: "scoreChange", Value: g.GetScoreChange()},
		{Key: "updatedAt", Value: time.Now()},
	}

	model := mongo.NewReplaceOneModel()
	model.SetFilter(filter)
	model.SetReplacement(replacement)
	model.SetUpsert(true)

	return model
}
