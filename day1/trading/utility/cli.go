package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//cli arguments
	args := os.Args
	result, err := configureDb(args)
	if err == nil {
		for index, value := range result {
			fmt.Println(index, value)
		}
	} else {
		fmt.Println(err)
	}

}

//multiple returns
func configureDb(properties []string) ([]string, error) {
	var dbProps = make([]string, 3)
	port, err := strconv.Atoi(properties[2])
	if err != nil {
		return nil, err
	}

	if port < 1027 {
		return nil, err
	} else {
		for index, value := range properties {
			dbProps[index] = value
		}
		return dbProps, nil
	}

}
