package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const baseUrl = "https://api.mozambiquehe.re/bridge?version=%s&platform=%s&uid=%s&auth=%s"

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

func buildPlayersURL(ids []string) string {
	uids := strings.Join(ids, ",")

	return fmt.Sprintf(baseUrl, apiVersion, platform, uids, apiKey)
}

func executeRequest(url string) (*http.Response, error) {
	return http.Get(url)
}

func getRequestBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
