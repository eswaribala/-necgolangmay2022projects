package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"necws/day3/dto"
)

func DBHelper() *sql.DB {
	db, err := sql.Open("mysql",
		"root:vignesh@(127.0.0.1:3306)/necdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateCurrency(Code string, Value int32, Symbol string) (int64, error) {

	db := DBHelper()
	defer db.Close()
	queryString := "Insert into Currency (Code,Value,Symbol) values(?,?,?)"
	result, err := db.Exec(queryString, Code, Value, Symbol)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	}
	return result.RowsAffected()

}

func ViewCurrencies() {

	db := DBHelper()
	defer db.Close()
	queryString := "Select * from Currency;"
	var currencyResponse dto.CurrencyResponse
	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal("Error occurred while saving...", err)
	} else {
		for rows.Next() {
			rows.Scan(&currencyResponse.CurrencyCode,
				&currencyResponse.CurrencyValue, &currencyResponse.CurrencySymbol)
			fmt.Printf("%v\n", currencyResponse)
		}

	}

}
