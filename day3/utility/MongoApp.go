package main

import (
	"fmt"
	"necws/day3/dao"
)

func main() {

	_, err := dao.MongoDbHelper()
	if err == nil {
		fmt.Println("Mongodb conn ready...")
	}

}
