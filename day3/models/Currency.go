package models

import "necws/day3/dao"

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

func (Curr *Currency) Create() (int64, error) {
	result, err := dao.CreateCurrency(Curr.CurrencyCode,
		Curr.CurrencyValue, Curr.CurrencySymbol)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (Curr *Currency) Edit(value int32) *Currency {
	Curr.CurrencyValue = value
	return Curr

}

func (Curr *Currency) Delete() {

}
