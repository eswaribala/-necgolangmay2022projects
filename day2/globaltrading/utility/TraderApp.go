package main

import (
	"fmt"
	"math/rand"
	"necws/day2/globaltrading/models"
	"reflect"
	"strconv"
)

var customers [2]models.Customer

func createCustomer() {
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
}

func viewCustomers() {
	//view the customer details

	for index, value := range customers {
		fmt.Printf("%d=>%v\n", index, value)
	}
}
func main() {

	createCustomer()
	viewCustomers()

	var customerObj models.Customer
	//type checking
	if reflect.TypeOf(customerObj).Name() == "Customer" {
		fmt.Println(reflect.TypeOf(customerObj).Name())
	}

	fmt.Printf("Type%T", customerObj)
}
