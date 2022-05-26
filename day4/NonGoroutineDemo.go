package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
	}

	for _, link := range links {
		accessLink(link)
	}
}

func accessLink(link string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error Occurred", err)
	} else {
		fmt.Println("Visiting", link)
	}

}
