package main

import "fmt"

func superb() {
	fmt.Println("Ahhaaashfdlhsl")
}

//catch block
func RecoverAccount() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func ValidateAccount(accountNo int32, password string) {
	defer RecoverAccount()
	defer superb() //executed at the end of the function
	if accountNo < 0 {
		panic("Account No should be 18 digit number") //try block
	}

	if len(password) == 0 {
		panic("Password cannot be null") //try block
	}

	fmt.Printf("%d,%s\n", accountNo, password)
	fmt.Println("Validate function return Normal")

}

func main() {
	defer fmt.Println("Running at the end from main")
	ValidateAccount(45454, "")

}
