package main

import (
	"fmt"
	"necws/day3/dao"
)

func main() {
	dbConn := dao.DBHelper()

	if dbConn != nil {
		fmt.Println("Mysql Conn ready")
	}

}
