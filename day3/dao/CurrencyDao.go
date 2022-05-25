package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
