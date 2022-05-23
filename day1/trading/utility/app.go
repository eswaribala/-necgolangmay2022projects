package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//initialize random
	rand.Seed(100000)

	//trading related info

	//customer
	var (
		firstName, lastName, dob, country string
		customerId                        int64
	)
	const F = "%T=%v\n"
	/*
		firstName = "Parameswari"
		lastName = "Bala"
		dob = "1970-12-02"
		country = "India"
		customerId = 2745763647
	*/
	//read the values
	//fmt.Println("Enter Customer Id")
	//fmt.Scanln(&customerId)

	customerId = rand.Int63()
	fmt.Println("Enter First Name")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name")
	fmt.Scanln(&lastName)

	fmt.Println("Enter DOB")
	fmt.Scanln(&dob)
	fmt.Println("Enter Country")
	fmt.Scanln(&country)
	fmt.Printf(F, firstName, firstName)
	fmt.Printf("Last Name%s\n", lastName)
	fmt.Printf("DOB %s\n Country%s\n", dob, country)
	fmt.Printf("CustomerId %d\n", customerId)
}
