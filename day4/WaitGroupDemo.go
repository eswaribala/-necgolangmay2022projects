package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
	}

	wg.Add(len(links))
	for _, link := range links {
		go accessLinkV4(link)
	}
	wg.Wait()

	fmt.Println("Main completes....")
}
func accessLinkV4(link string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error Occurred", err)
	} else {
		fmt.Println("Visiting", link)
	}

	wg.Done()
}
