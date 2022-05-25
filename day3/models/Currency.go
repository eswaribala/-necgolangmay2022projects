package models

type Currency struct {
	CurrencyCode   string
	CurrencyValue  int32
	CurrencySymbol string
}

//Embedded

type Transaction struct {
	TransactionId     int64
	TransactionAmount int64
	Currency
}

type ICurrency interface {
	Create()
	Edit()
	Delete()
}

func (Curr *Currency) Create() {

}
func (Curr *Currency) Edit(value int32) *Currency {
	Curr.CurrencyValue = value
	return Curr

}

func (Curr *Currency) Delete() {

}
