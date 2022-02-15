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
	data := fetchData(usernames)
	players := deserializeData(data, len(usernames))

	updatePlayersData(players)

	log.Printf("Found users %+v\n", players)
}

func fetchData(usernames []string) []byte {
	log.Println("Fetching player info for [", strings.Join(usernames, ","), "]...")
	data, err := fetchPlayersData(usernames)

	if err != nil {
		log.Fatalln("Error fetching players!")
		log.Fatalln(err)

		os.Exit(1)
	}

	log.Println("Fetched player data!")
	return data
}

func deserializeData(data []byte, nrUsers int) []PlayerData {
	log.Println("Deserializing player data...")
	users, err := deserializePlayers(data, nrUsers)

	if err != nil {
		log.Fatalln("Error deserializing player data!")
		log.Fatalln(err)

		os.Exit(1)
	}

	log.Println("Deserialized players!")
	log.Println("Number of players found: ", len(users))
	return users
}

func updatePlayersData(players []PlayerData) {
	log.Println("Updating players data...")
	err := updatePlayers(players)

	if err != nil {
		log.Fatalln("Error updating player data!")
		log.Fatalln(err)

		os.Exit(1)
	}

	log.Println("Updated players data!")
}
