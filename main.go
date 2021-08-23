package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	color.White.Println(" ~ HTTP REQUEST INIT ~ ")
	REQUEST_LINK := getConnectionString()
	NEW_PARAM := getRequest()
	fmt.Println(REQUEST_LINK)
	fmt.Println(NEW_PARAM)

	makeRequest(NEW_PARAM)
}

func getConnectionString() string {
	var URL string
	fmt.Print("Enter URL: ")
	fmt.Scanln(&URL)

	return URL

}

func getRequest() string {
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/ping", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	query := req.URL.Query()
	query.Add("api_key", "KEY_GOES_HERE")
	req.URL.RawQuery = query.Encode()

	return req.URL.String()
}

// TODO: Add HTTP Method
func makeRequest(URL string) {
	response, error := http.Get(URL)
	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	body, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Fatal(error)
	}

	color.Red.Println("HTTP RESPONSE FROM URL")
	fmt.Println(string(body))

}
