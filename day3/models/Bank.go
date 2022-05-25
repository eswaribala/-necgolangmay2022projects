package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"necws/day3/dao"
	"necws/day3/dto"
)

type Bank struct {
	ID         primitive.ObjectID `bson:"_id"`
	IFSCCode   string             `bson:"IFSC_Code"`
	BankName   string             `bson:"Bank_Name"`
	BranchName string             `bson:"Branch_Name"`
}

type BankSorter []Bank

type IBank interface {
	Create()
	Edit()
	Delete()
}

//total no of elements
func (bankSorter BankSorter) Len() int {
	return len(bankSorter)
}

//swap
func (bankSorter BankSorter) Swap(i, j int) {
	bankSorter[i], bankSorter[j] = bankSorter[j], bankSorter[i]
}

//attribue/property
func (bankSorter BankSorter) Less(i, j int) bool {
	return bankSorter[i].BankName < bankSorter[j].BankName
}
func (bank *Bank) Create() {
	var bankRequest dto.BankRequest = dto.BankRequest{
		bank.ID, bank.IFSCCode, bank.BankName,
		bank.BranchName,
	}
	dao.CreateBank(bankRequest)
}

func Edit() {

}

func Delete() {

}
