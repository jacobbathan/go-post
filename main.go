package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ~ HTTP REQUEST INIT ~ ")
	REQUEST_LINK := getConnectionString()

	fmt.Println(REQUEST_LINK)
}

func getConnectionString() string {
	var URL string
	fmt.Print("Enter URL: ")
	fmt.Scanln(&URL)

	return URL

}
