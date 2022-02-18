package apigateway

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var apiKey = os.Getenv("API_KEY")

func executeRequest(url string) *http.Response {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("base_client#executeRequest - Error executing request", err)
	}

	if resp.StatusCode >= 400 {
		log.Fatalln("base_client#executeRequest - Request returned status code", resp.StatusCode)
	}

	return resp
}

func getResponseBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("base_client#getRequestBody - Error fetching request body", err)
	}

	return body
}
