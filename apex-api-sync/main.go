package main

import (
	"log"

	"apex-api-sync/apigateway"
	"apex-api-sync/models"
	"apex-api-sync/repository"
)

func main() {
	log.Println("Updating player data...")

	ids := fetchPlayerIDs()

	log.Println("Updating player data for ids", ids)

	players := fetchData(ids)

	updatePlayersData(players)

	log.Println("Updated player data for ids", ids)

	log.Println("Updating games...")

	games := make([]models.GameData, 0)

	for _, id := range ids {
		playerGames := apigateway.GetGames(id)

		games = append(games, playerGames...)
	}

	repository.UpdateGames(games)

	log.Println("Updated games!")
}

func fetchPlayerIDs() []string {
	log.Println("Fetching player IDs...")
	return repository.FetchPlayerIDs()
}

func fetchData(ids []string) []models.PlayerData {
	log.Println("Fetching player data...")
	data := apigateway.GetPlayers(ids)
	log.Println("Fetched player data!")
	return data
}

func updatePlayersData(players []models.PlayerData) {
	log.Println("Updating players data...")
	repository.UpdatePlayers(players)
}
