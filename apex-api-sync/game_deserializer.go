package main

import (
	"apex-api-sync/models"
	"encoding/json"
)

func deserializeGames(data []byte) ([]models.GameData, error) {
	games := make([]models.GameData, 0)

	err := json.Unmarshal(data, &games)

	return games, err
}
