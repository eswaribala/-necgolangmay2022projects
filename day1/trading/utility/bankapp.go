package main

import (
	"fmt"
	models "necws/day1/trading/models"
)

func main() {

	var customer models.Customer

	customer = models.Customer{48358, "Parameswari",
		"Bala", "param@gmail.com", true}

	fmt.Printf("Customer=%v", customer)

}
