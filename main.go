package main

import (
	"log"
	"os"
	"strings"
)

var usernames = []string{
	"Leoniverse",
	"monkasSFrEE",
}

func main() {
	log.Println("Updating player data for [", strings.Join(usernames, ","), "]...")

	data := fetchData(usernames)
	players := deserializeData(data, len(usernames))

	updatePlayersData(players)
	
	log.Println("Updated player data for [", strings.Join(usernames, ","), "]!")
}

func fetchData(usernames []string) []byte {
	log.Println("Fetching player data...")
	data, err := fetchPlayersData(usernames)

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
