package repository

import (
	"apex-api-sync/models"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var playersDB = os.Getenv("MONGO_PLAYERS_DATABASE")
var playersCol = os.Getenv("MONGO_PLAYERS_COLLECTION")

var playersCollection = getCollection(playersDB, playersCol)

func UpdatePlayers(players []models.PlayerData) {
	models := getPlayerWriteModels(players)

	options := getBulkWriteOptions()

	_, err := playersCollection.BulkWrite(context.Background(), models, options)

	if err != nil {
		log.Fatalln("games_repository#UpdateGames - Error updating games", err)
	}
}

func FetchPlayerIDs() []string {
	res, err := playersCollection.Distinct(context.Background(), "_id", bson.D{})

	if err != nil {
		log.Fatalln("players_repository#FetchUIDs - Error fetching player ids", err)
	}

	return extractIds(res)
}

func getPlayerWriteModels(players []models.PlayerData) []mongo.WriteModel {
	models := make([]mongo.WriteModel, 0)

	for _, p := range players {
		models = append(models, playerDataToWriteModel(p))
	}

	return models
}

func extractIds(results []interface{}) []string {
	ids := make([]string, 0)

	for _, uid := range results {
		ids = append(ids, fmt.Sprintf("%v", uid))
	}

	return ids
}
