package apigateway

import (
	"apex-api-sync/models"
	"fmt"
)

const getGamesURL = "https://api.mozambiquehe.re/games?auth=%s&uid=%s"

func GetGames(playerId string) []models.GameData {
	url := buildGetGamesURL(playerId)

	resp := executeRequest(url)

	body := getResponseBody(resp)

	return deserializeGames(body)
}

func buildGetGamesURL(playerId string) string {
	return fmt.Sprintf(getGamesURL, apiKey, playerId)
}
