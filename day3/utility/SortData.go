package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var data = make([]int, 100)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100000)
	}
	for _, value := range data {
		fmt.Printf("%d\n", value)
	}

}
