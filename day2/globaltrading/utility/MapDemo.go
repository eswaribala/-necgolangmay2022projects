package main

import (
	"fmt"
	"strconv"
)

func main() {

	//declare map
	var countries = make(map[int]string)

	//assign values
	for i := 0; i < 5; i++ {
		countries[i] = "Country" + strconv.Itoa(i)
	}

	//read values
	for key, value := range countries {
		fmt.Printf("%d=>%s\n", key, value)
	}

}
