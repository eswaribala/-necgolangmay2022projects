package main

import "fmt"

func main() {
	//string is immutable
	bankName := "JP Morgan"
	//bankName[1] = 'p'
	fmt.Println(bankName)

	//rune
	traderBankName := []rune("JP Morgan")
	traderBankName[1] = 'p'
	convBankName := string(traderBankName)
	fmt.Println(convBankName)

	//initalizing variable

	tradingDate := "2022-05-20" //initializing, dynamically infers
	tradingDate = "2022-05-21"  //reassign the value
	fmt.Println(tradingDate)

}
