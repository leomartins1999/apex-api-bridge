package apigateway

import (
	"apex-api-sync/models"
	"fmt"
	"strings"
)

const getPlayersURL = "https://api.mozambiquehe.re/bridge?version=5&platform=PC&auth=%s&uid=%s"

func GetPlayers(ids []string) []models.PlayerData {
	url := buildGetPlayersURL(ids)

	resp := executeRequest(url)

	body := getResponseBody(resp)

	return deserializePlayers(body, len(ids))
}

func buildGetPlayersURL(ids []string) string {
	uids := strings.Join(ids, ",")

	return fmt.Sprintf(getPlayersURL, apiKey, uids)
}
