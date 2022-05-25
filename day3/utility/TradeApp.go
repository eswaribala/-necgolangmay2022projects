package main

import "necws/day3/dao"

func main() {
	/*
		Curr := models.Currency{"SGD",
			45, "SGD"}
		result, error := Curr.Create()
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Printf("Records inserted %d", result)
		}
	*/
	//var curr models.Currency = models.Currency{}
	//curr.View()
	dao.ViewCurrencies()

}
