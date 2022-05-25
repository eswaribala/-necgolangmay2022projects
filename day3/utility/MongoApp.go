package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"necws/day3/models"
)

func main() {

	var bank models.Bank = models.Bank{primitive.ObjectID{},
		"HSBC000111",
		"HSBC", "Mylapore",
	}
	bank.Create()
}
