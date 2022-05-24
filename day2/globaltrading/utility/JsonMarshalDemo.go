package main

import (
	"encoding/json"
	"fmt"
	"necws/day2/globaltrading/models"
)

func main() {

	var customerObj = models.Customer{
		13423423, "Parameswari", "India",
		models.Address{}, models.Bank{},
	}

	//marshalling
	customer, err := json.Marshal(customerObj)

	if err == nil {
		fmt.Println(string(customer))
	}
}
