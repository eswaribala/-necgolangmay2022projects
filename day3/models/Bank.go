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

type IBank interface {
	Create()
	Edit()
	Delete()
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
