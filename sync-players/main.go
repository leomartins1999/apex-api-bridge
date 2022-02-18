package main

import (
	"log"

	"apex-api-sync/apigateway"
	"apex-api-sync/repository"
)

func main() {
	log.Println("### Updating players... ###")

	log.Println("Fetching player ids...")
	ids := repository.FetchPlayerIDs()
	log.Println("Number of players found:", len(ids))

	log.Println("Fetching player data...")
	players := apigateway.GetPlayers(ids)

	log.Println("Updating player data...")
	repository.UpdatePlayers(players)

	log.Println("### Updated players! ###")
}
