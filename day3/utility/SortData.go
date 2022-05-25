package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"necws/day3/models"
	"sort"
)

func main() {
	var data = make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100000)
	}
	//	fmt.Println("Before Sorting.....")
	/*
		for _, value := range data {
			fmt.Printf("%d\n", value)
		}
	*///sort the data
	sort.Ints(data)
	//fmt.Println("After Sorting.....")
	/*
		for _, value := range data {
			fmt.Printf("%d\n", value)
		}
	*/
	bank1 := models.Bank{primitive.ObjectID{}, "H00001",
		"HSBC", "Mylapore"}
	bank2 := models.Bank{primitive.ObjectID{}, "S00001",
		"SBI", "Mylapore"}
	bank3 := models.Bank{primitive.ObjectID{}, "A00001",
		"AXIS", "Mylapore"}
	bank4 := models.Bank{primitive.ObjectID{}, "I00001",
		"ICICI", "Mylapore"}

	var banksorters = make(models.BankSorter, 5)
	banksorters[0] = bank1
	banksorters[1] = bank2
	banksorters[2] = bank3
	banksorters[3] = bank4
	sort.Sort(banksorters)

	for _, value := range banksorters {
		fmt.Printf("%v\n", value)
	}

}
