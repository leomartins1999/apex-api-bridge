package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"apex-api-sync/models"
)

const playersURL = "https://api.mozambiquehe.re/bridge?version=%s&platform=%s&uid=%s&auth=%s"
const gamesURL = "https://api.mozambiquehe.re/games?auth=%s&uid=%s"

const apiVersion = "5"
const platform = "PC"

var apiKey = os.Getenv("API_KEY")

func fetchPlayersData(ids []string) ([]byte, error) {
	url := buildPlayersURL(ids)

	resp, err := executeRequest(url)
	if err != nil {
		return []byte{}, err
	}

	return getRequestBody(resp)
}

func fetchGames(playerId string) ([]models.GameData, error) {
	url := buildGamesURL(playerId)

	resp, err := executeRequest(url)
	if err != nil {
		return []models.GameData{}, err
	}

	body, err := getRequestBody(resp)
	if err != nil {
		return []models.GameData{}, err
	}

	return deserializeGames(body)
}

func buildPlayersURL(ids []string) string {
	uids := strings.Join(ids, ",")

	return fmt.Sprintf(playersURL, apiVersion, platform, uids, apiKey)
}

func buildGamesURL(playerId string) string {
	return fmt.Sprintf(gamesURL, apiKey, playerId)
}

func executeRequest(url string) (*http.Response, error) {
	return http.Get(url)
}

func getRequestBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
