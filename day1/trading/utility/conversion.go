package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//global variable
var fromCurrency string = "INR"

//created array
var currencies = []string{"INR", "USD"}

var ids = make([]int32, 100)

func main() {

	//var stradingLimit string
	//updating
	var currencies = append(currencies, "SAR")

	for index, currency := range currencies {
		//currencies = append(currencies, "DHIRAM")
		fmt.Printf("Index %d=>%s\n", index, currency)
	}

	//slicing
	var slicedCurrencies = currencies[1:2]
	fmt.Println("Sliced Currencies")
	for index, currency := range slicedCurrencies {
		//currencies = append(currencies, "DHIRAM")
		fmt.Printf("Index %d=>%s\n", index, currency)
	}

	//generated ids

	for index, value := range generateCurrencyId() {
		fmt.Printf("Index %d=>%d\n", index, value)
	}

	/*fmt.Println("Enter Trading Limit")
	fmt.Scanln(&stradingLimit)
	//dynamic data type
	var tradingLimit, err = strconv.Atoi(stradingLimit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tradingLimit)
	}
	*/
	var sexchangeRate string
	fmt.Println("Enter Exchange Rate")
	fmt.Scanln(&sexchangeRate)

	var exchangeRate, error = strconv.ParseInt(sexchangeRate, 10, 8)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(exchangeRate)
	}

	currencyConversion("USD", exchangeRate, 20000000)

}

func currencyConversion(toCurrency string, currencyRate int64, amount int64) {
	//local variable
	var convertedAmount int64
	if fromCurrency == "INR" {
		convertedAmount = amount / currencyRate
		fmt.Printf("Converted Amount %d %s", convertedAmount, toCurrency)

	}

}

func generateCurrencyId() []int32 {

	for pos := range ids {
		ids[pos] = rand.Int31n(1000000)
	}

	return ids

}
