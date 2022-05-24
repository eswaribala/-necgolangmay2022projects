package main

import (
	"necws/day2/globaltrading/models"
)

func main() {

	var customerObj models.Customer //struct
	var pcustomerObj *models.Customer
	var iCustomer models.ICustomer //interface
	customerObj = models.Customer{
		2384585, "Shashank", "India",
		models.Address{}, models.Bank{},
	} //object
	customerObj.View()
	pcustomerObj = &customerObj
	pcustomerObj.View()
	iCustomer = &customerObj
	iCustomer.View()

}
