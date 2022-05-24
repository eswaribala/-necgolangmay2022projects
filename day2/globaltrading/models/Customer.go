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

//call by value
//updating the object
//receiveobject method name method param return type
func (c Customer) ChangeCountryName(countryName string) Customer {
	//modify
	c.Country = countryName
	return c
}

//call by reference
func (c *Customer) ChangeName(Name string) *Customer {
	//modify
	c.Name = Name
	return c
}
