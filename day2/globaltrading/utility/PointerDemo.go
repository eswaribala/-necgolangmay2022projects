package main

import "fmt"

func main() {

	var traderName string = "Akshay"

	var ptraderName *string = &traderName

	fmt.Println(traderName)
	//to read the value of pointed variable
	fmt.Println(*ptraderName)
}
