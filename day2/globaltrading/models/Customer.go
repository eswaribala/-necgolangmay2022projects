package models

type Bank struct {
	AccountNo  int64
	IFSCCode   string
	BranchName string
}

type Address struct {
	DoorNo     int16
	StreetName string
	City       string
	State      string
}

type Customer struct {
	CustomerId  int64
	Name        string
	Country     string
	AddressData Address
	BankDetail  Bank
}
