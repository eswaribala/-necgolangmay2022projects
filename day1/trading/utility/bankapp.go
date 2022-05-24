package main

import (
	"fmt"
	models "necws/day1/trading/models"
)

func main() {

	var customer models.Customer

	customer = models.Customer{48358, "Parameswari",
		"Bala", "param@gmail.com", true,
		models.Date{2, 12, 1970}}

	fmt.Printf("Customer=%v", customer)

}
