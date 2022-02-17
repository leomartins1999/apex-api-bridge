package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Updating player data...")

	ids := fetchPlayerIDs()

	log.Println("Updating player data for ids", ids)

	data := fetchData(ids)
	players := deserializeData(data, len(ids))

	updatePlayersData(players)

	log.Println("Updated player data for ids", ids)
}

func fetchPlayerIDs() []string {
	log.Println("Fetching player IDs...")
	ids, err := fetchUIDs()

	if err != nil {
		log.Fatalln("Error fetching player IDs !", err)
		os.Exit(1)
	}

	log.Println("Fetched player IDs!")
	return ids
}

func fetchData(ids []string) []byte {
	log.Println("Fetching player data...")
	data, err := fetchPlayersData(ids)

	if err != nil {
		log.Fatalln("Error fetching players!", err)
		os.Exit(1)
	}

	log.Println("Fetched player data!")
	return data
}

func deserializeData(data []byte, nrUsers int) []PlayerData {
	log.Println("Deserializing player data...")
	users, err := deserializePlayers(data, nrUsers)

	if err != nil {
		log.Fatalln("Error deserializing player data!", err)
		os.Exit(1)
	}

	log.Println("Deserialized players!")
	return users
}

func updatePlayersData(players []PlayerData) {
	log.Println("Updating players data...")
	err := updatePlayers(players)

	if err != nil {
		log.Fatalln("Error updating player data! ", err)
		os.Exit(1)
	}

	log.Println("Updated players data!")
}
