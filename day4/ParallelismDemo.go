package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func listenMusic(wg *sync.WaitGroup) {

	fmt.Println("Listening Music....")
	wg.Done()
}

func ridingCar(wg *sync.WaitGroup) {
	fmt.Println("Riding Car....")
	wg.Done()
}

func watchingMovie(wg *sync.WaitGroup) {
	fmt.Println("Watching TV....")
	wg.Done()
}

func eatingPopCorn(wg *sync.WaitGroup) {
	fmt.Println("Eating Popcorn....")
	wg.Done()
}

var listOfTasks = []func(wg *sync.WaitGroup){
	listenMusic, ridingCar, watchingMovie, eatingPopCorn,
}

func main() {

	fmt.Println(len(listOfTasks))
	wg2.Add(len(listOfTasks))
	for _, task := range listOfTasks {

		go task(&wg2)
	}
	fmt.Println("All the tasks done")
	wg2.Wait()
}
