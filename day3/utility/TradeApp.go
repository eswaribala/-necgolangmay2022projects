package main

import "necws/day3/models"

func main() {
	/*
		Curr := models.Currency{"MALE",
			30, "MALE"}
		result, error := Curr.Create()
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Printf("Records inserted %d\n", result)
		}
	*/
	var curr models.Currency = models.Currency{}
	curr.View()

}
