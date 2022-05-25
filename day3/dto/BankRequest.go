package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type BankRequest struct {
	ID         primitive.ObjectID `bson:"_id"`
	IFSCCode   string             `bson:"IFSC_Code"`
	BankName   string             `bson:"Bank_Name"`
	BranchName string             `bson:"Branch_Name"`
}
