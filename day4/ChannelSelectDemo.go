package main

import (
	"fmt"
	"strings"
)

func main() {
	//Economic Trading
	et := make(chan string)
	//Non Economic Trading
	net := make(chan string)
	//Delta calculation
	dc := make(chan string)

	//anonymous function
	go func(message string, dc chan string) {
		if strings.ContainsAny(message, "Delta") {
			dc <- "Delta Ready"
		} else {
			dc <- "No Delta Calculations"
		}

	}("ComputeDelta", dc)

	go ecoTrading("Compute ET", et)
	go nonEcoTrading("Compute Leg1 and Leg2", net)

	var decision string
	fmt.Println("Do you want to close Trading")
	fmt.Scanln(&decision)

}

func ecoTrading(message string, et chan string) {

	if len(message) == 0 {
		et <- "Message cannot be empty"
	} else {
		et <- "Eco Trading Done"
	}

}

func nonEcoTrading(message string, net chan string) {

	if strings.HasPrefix(message, "Leg1") {
		net <- "Non Economic Trading Done..."
	} else {
		net <- "Not Done..."
	}

}
