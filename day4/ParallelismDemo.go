package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var wg2 sync.WaitGroup

func listenMusic(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Listening Music...." + strconv.Itoa(i))
	}

}

func ridingCar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Riding Car...." + strconv.Itoa(i))
	}

}

func watchingMovie(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Watching TV...." + strconv.Itoa(i))
	}

}

func eatingPopCorn(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Eating pop corn...." + strconv.Itoa(i))
	}

}

var listOfTasks = []func(wg *sync.WaitGroup){
	listenMusic, ridingCar, watchingMovie, eatingPopCorn,
}

func main() {
	runtime.GOMAXPROCS(7)
	fmt.Println(len(listOfTasks))
	wg2.Add(len(listOfTasks))
	for _, task := range listOfTasks {

		go task(&wg2)
	}

	wg2.Wait()
}
