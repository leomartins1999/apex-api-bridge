package main

import (
	"log"

	"apex-api-sync/apigateway"
	"apex-api-sync/deserializers"
	"apex-api-sync/models"
	"apex-api-sync/repositories"
)

func main() {
	log.Println("Updating player data...")

	ids := fetchPlayerIDs()

	log.Println("Updating player data for ids", ids)

	data := fetchData(ids)
	players := deserializeData(data, len(ids))

	updatePlayersData(players)

	log.Println("Updated player data for ids", ids)

	log.Println("Updating games...")

	games := make([]models.GameData, 0)

	for _, id := range ids {
		playerGames, err := apigateway.FetchGames(id)

		if err != nil {
			log.Fatalln("Error fetching games for player ", id, err)
		}

		games = append(games, playerGames...)
	}

	repositories.UpdateGames(games)

	log.Println("Updated games!")
}

func fetchPlayerIDs() []string {
	log.Println("Fetching player IDs...")
	ids, err := repositories.FetchUIDs()

	if err != nil {
		log.Fatalln("Error fetching player IDs !", err)
	}

	log.Println("Fetched player IDs!")
	return ids
}

func fetchData(ids []string) []byte {
	log.Println("Fetching player data...")
	data, err := apigateway.FetchPlayersData(ids)

	if err != nil {
		log.Fatalln("Error fetching players!", err)
	}

	log.Println("Fetched player data!")
	return data
}

func deserializeData(data []byte, nrUsers int) []models.PlayerData {
	log.Println("Deserializing player data...")
	users, err := deserializers.DeserializePlayers(data, nrUsers)

	if err != nil {
		log.Fatalln("Error deserializing player data!", err)
	}

	log.Println("Deserialized players!")
	return users
}

func updatePlayersData(players []models.PlayerData) {
	log.Println("Updating players data...")
	err := repositories.UpdatePlayers(players)

	if err != nil {
		log.Fatalln("Error updating player data! ", err)
	}

	log.Println("Updated players data!")
}
