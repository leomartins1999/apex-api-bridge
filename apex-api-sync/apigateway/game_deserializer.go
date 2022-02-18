package apigateway

import (
	"apex-api-sync/models"
	"encoding/json"
	"log"
)

func deserializeGames(data []byte) []models.GameData {
	games := make([]models.GameData, 0)

	err := json.Unmarshal(data, &games)

	if err != nil {
		log.Fatalln("game_deserializer#deserializeGames - Error deserializing games. Data: ", string(data), "Error: ", err)
	}

	return games
}
