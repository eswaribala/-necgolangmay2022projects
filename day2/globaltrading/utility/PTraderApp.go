package main

import (
	"fmt"
	"necws/day2/globaltrading/models"
)

func main() {

	var customerObj models.Customer
	var pcustomerObj *models.Customer
	customerObj = models.Customer{28487, "Parameswari",
		"India", models.Address{}, models.Bank{}}

	pcustomerObj = &customerObj
	pcustomerObj.ChangeName("EswariBala")
	fmt.Printf("%v\n", *pcustomerObj)
	fmt.Printf("%v\n", customerObj)
}
