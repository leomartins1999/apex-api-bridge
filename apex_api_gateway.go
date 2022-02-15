package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const baseUrl = "https://api.mozambiquehe.re/bridge?version=%s&platform=%s&player=%s&auth=%s"

const apiVersion = "5"
const platform = "PC"

var apiKey = os.Getenv("API_KEY")

func fetchPlayersData(usernames []string) ([]byte, error) {
	url := buildPlayersURL(usernames)

	resp, err := executeRequest(url)
	if err != nil {
		return []byte{}, err
	}

	return getRequestBody(resp)
}

func buildPlayersURL(usernames []string) string {
	players := strings.Join(usernames, ",")

	return fmt.Sprintf(baseUrl, apiVersion, platform, players, apiKey)
}

func executeRequest(url string) (*http.Response, error) {
	return http.Get(url)
}

func getRequestBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
