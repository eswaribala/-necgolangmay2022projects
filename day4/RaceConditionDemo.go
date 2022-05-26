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
	go deposit(10000)
	wg1.Wait()
}

func deposit(amount int) {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		availableBalance := accountBalance
		accountBalance += amount
		fmt.Printf("Available Balance%d\t", availableBalance)
		fmt.Printf("Account Balance %d\n", accountBalance)

	}
}
