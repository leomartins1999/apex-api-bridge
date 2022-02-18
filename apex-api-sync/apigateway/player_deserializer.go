package apigateway

import (
	"apex-api-sync/models"
	"encoding/json"
	"log"
)

func deserializePlayers(data []byte, nrPlayers int) []models.PlayerData {
	var result []models.PlayerData
	var err error

	if nrPlayers == 1 {
		result, err = deserializeSinglePlayer(data)
	} else {
		result, err = deserializeMultiplePlayers(data)
	}

	if err != nil {
		log.Fatalln("player_deserializer#deserializePlayers - Error deserializing players. Data: ", string(data), "Error: ", err)
	}

	return result
}

func deserializeSinglePlayer(data []byte) ([]models.PlayerData, error) {
	var p models.PlayerData

	err := json.Unmarshal(data, &p)

	return []models.PlayerData{p}, err
}

func deserializeMultiplePlayers(data []byte) ([]models.PlayerData, error) {
	players := make([]models.PlayerData, 0)

	err := json.Unmarshal(data, &players)

	return players, err
}
