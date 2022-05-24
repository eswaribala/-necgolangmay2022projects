package models

type Date struct {
	Day   int16
	Month int16
	Year  int16
}

type Customer struct {
	CustomerId int64
	FirstName  string
	LastName   string
	Email      string
	Status     bool
	DOB        Date
}
