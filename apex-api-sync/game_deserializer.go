package main

import "encoding/json"

func deserializeGames(data []byte) ([]GameData, error) {
	games := make([]GameData, 0)

	err := json.Unmarshal(data, &games)

	return games, err
}
