package models

type Bank struct {
	IFSCCode   string
	BankName   string
	BranchName string
}

type IBank interface {
	Create()
	Edit()
	Delete()
}

func Create() {

}

func Edit() {

}

func Delete() {

}