package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup
var accountBalance int = 5000

func main() {
	wg1.Add(2)
	go deposit(10000)
	go viewBalance()
	wg1.Wait()
}

func deposit(amount int) {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		//writing
		accountBalance += amount
		fmt.Printf("Account Balance %d\n", accountBalance)

	}
}

func viewBalance() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		//reading
		availableBalance := accountBalance
		fmt.Printf("Available Balance%d\n", availableBalance)

	}
}
