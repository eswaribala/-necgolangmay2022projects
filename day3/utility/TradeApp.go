package main

import (
	"fmt"
	"necws/day3/models"
)

func main() {
	transaction := models.Transaction{28487,
		25486583,
		models.Currency{"INR", 78,
			"Rupay"},
	}
	transaction.Edit(80)

	fmt.Printf("%v", transaction)
}
