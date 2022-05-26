package main

import (
	"fmt"
	"net/http"
)

func main() {
	//main routine starts
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
	}
	//create the channel
	c := make(chan string)

	for _, link := range links {
		//child routine starts
		//passing channel as parameter
		go accessLinkV1(link, c)
	}
	fmt.Println(<-c)

	//main routing executing
	fmt.Println("Completed all routine calls..... ")
}

//receiving channel parameter
func accessLinkV1(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error Occurred", err)
		//writing message in to channel
		c <- "Error occurred"
	} else {
		fmt.Println("Visiting", link)
		//writing message in to channel
		c <- "Link up"
	}

}
