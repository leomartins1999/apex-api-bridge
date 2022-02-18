package main

import (
	"log"

	"apex-sync-commons/apigateway"
	"apex-sync-commons/models"
	"apex-sync-commons/repository"
)

func main() {
	log.Println("### Updating games... ###")

	log.Println("Fetching player ids...")
	ids := repository.FetchPlayerIDs()
	log.Println("Number of players found:", len(ids))

	log.Println("Fetching games for found players...")
	games := getGames(ids)

	log.Println("Updating games...")
	repository.UpdateGames(games)

	log.Println("### Updated games! ###")
}

func getGames(playerIds []string) []models.GameData {
	games := make([]models.GameData, 0)

	for _, playerId := range playerIds {
		log.Println("Fetching games for player ", playerId)
		playerGames := apigateway.GetGames(playerId)
		games = append(games, playerGames...)
	}

	return games
}
