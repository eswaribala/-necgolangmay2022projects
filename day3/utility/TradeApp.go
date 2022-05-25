package main

import (
	"fmt"
	"necws/day3/models"
)

func main() {

	Curr := models.Currency{"INR",
		75, "Rupay"}
	result, error := Curr.Create()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("Records inserted %d", result)
	}

}
