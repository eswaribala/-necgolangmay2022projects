package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
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
	fmt.Printf("No of Cores %d\n", runtime.NumCPU())
	fmt.Println(len(listOfTasks) + 1)
	t1 := time.Now()
	wg2.Add(len(listOfTasks))
	for _, task := range listOfTasks {

		go task(&wg2)
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	wg2.Wait()

}
