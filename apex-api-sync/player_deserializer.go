package main

import (
	"apex-api-sync/models"
	"encoding/json"
)

func deserializePlayers(data []byte, nrPlayers int) ([]models.PlayerData, error) {
	if nrPlayers == 1 {
		return deserializeSinglePlayer(data)
	} else {
		return deserializeMultiplePlayers(data)
	}
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
