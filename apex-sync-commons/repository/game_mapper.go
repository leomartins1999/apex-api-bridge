package repository

import (
	"apex-sync-commons/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
