package repository

import (
	"apex-sync-commons/models"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var gamesDB = os.Getenv("MONGO_GAMES_DATABASE")
var gamesCol = os.Getenv("MONGO_GAMES_COLLECTION")

var gamesCollection = getCollection(gamesDB, gamesCol)

func UpdateGames(games []models.GameData) {
	models := getGameWriteModels(games)

	options := getBulkWriteOptions()

	_, err := gamesCollection.BulkWrite(context.Background(), models, options)

	if err != nil {
		log.Fatalln("games_repository#UpdateGames - Error updating games", err)
	}
}

func getGameWriteModels(games []models.GameData) []mongo.WriteModel {
	models := make([]mongo.WriteModel, 0)

	for _, g := range games {
		models = append(models, gameDataToWriteModel(g))
	}

	return models
}
