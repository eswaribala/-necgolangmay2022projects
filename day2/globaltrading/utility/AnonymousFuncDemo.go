package main

import "fmt"

func computeInterest(capital float32, roi float32) func() float32 {
	return func() float32 {
		return capital * roi * 12
	}
}

func main() {

	calcInterest := computeInterest(2000000, 0.2)
	fmt.Println(calcInterest())

	//anonymous function
	result := func(capital float32, roi float32) float32 {
		return capital * roi * 12
	}(12349759, 0.4)
	fmt.Println(result)
}
