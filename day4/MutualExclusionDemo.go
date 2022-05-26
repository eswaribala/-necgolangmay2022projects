package main

import (
	"fmt"
	"sync"
)

var wgg sync.WaitGroup
var mutex sync.Mutex
var balance int = 0

func main() {

	wgg.Add(3)
	go counter()
	go counter()
	go counter()

	wgg.Wait()

}

func counter() {

	defer wgg.Done()
	mutex.Lock()
	for i := 0; i < 100; i++ {
		balance += i
		fmt.Println(balance)
	}
	mutex.Unlock()
}
