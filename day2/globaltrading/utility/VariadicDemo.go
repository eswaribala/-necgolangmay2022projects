package main

import (
	"errors"
	"fmt"
	"strconv"
)

//variadic function

func LIBORData(bankName string, roi ...int32) (string, error) {
	var sum int32
	if len(roi) == 0 {
		return "", errors.New("Interest Rate is Empty")
	} else {
		for _, value := range roi {
			sum += value
		}
		average := sum / int32(len(roi))
		fmt.Printf("Average%d", average)
		return "Average ROI is" + strconv.Itoa(int(average)), nil
	}

}

func main() {

	result, err := LIBORData("RBS", 12, 13, 4, 5, 13)

	if err == nil {
		fmt.Printf("%s\n", result)
	}

}
