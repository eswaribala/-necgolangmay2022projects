package main

import (
	"fmt"
	"math/rand"
	"necws/day2/globaltrading/models"
	"strconv"
)

func main() {

	var customers [2]models.Customer

	for i := 0; i < len(customers); i++ {
		customers[i] = models.Customer{rand.Int63n(1000000),
			"Customer" + strconv.Itoa(i), "India",
			models.Address{
				int16(rand.Int31()), "Rajaji st", "Chennai",
				"TamilNadu",
			},
			models.Bank{rand.Int63n(1000000),
				"SBI000174",
				"Adyar Branch",
			},
		}
	}

	//view the customer details

	for index, value := range customers {
		fmt.Printf("%d=>%v\n", index, value)
	}

}
