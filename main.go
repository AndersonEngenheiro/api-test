package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define the URL you want to request
	url := "https://economia.awesomeapi.com.br/last/USD-BRL,EUR-BRL,BTC-"

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending GET request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Print the response
	fmt.Printf("Response Body: %s\n", string(responseBody))
}
