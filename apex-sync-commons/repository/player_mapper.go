package repository

import (
	"apex-sync-commons/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
