package main

import (
	"encoding/json"
)

func deserializePlayers(data []byte, nrPlayers int) ([]PlayerData, error) {
	if nrPlayers == 1 {
		return deserializeSinglePlayer(data)
	} else {
		return deserializeMultiplePlayers(data)
	}
}

func deserializeSinglePlayer(data []byte) ([]PlayerData, error) {
	var p PlayerData

	err := json.Unmarshal(data, &p)

	return []PlayerData{p}, err
}

func deserializeMultiplePlayers(data []byte) ([]PlayerData, error) {
	players := make([]PlayerData, 0)

	err := json.Unmarshal(data, &players)

	return players, err
}
