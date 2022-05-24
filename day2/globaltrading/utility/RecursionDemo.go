package main

import "fmt"

func OTPCountDown(number int) {

	if number > 0 {
		fmt.Println(number)
		OTPCountDown(number - 1)
	} else {
		fmt.Println("Time Exceeded, Generate OTP")
	}

}

func main() {
	OTPCountDown(60)
}
