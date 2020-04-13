package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	time.Sleep(time.Second)
	for v := range c {
		fmt.Printf("worker %d received %d\n", id, v)
	}
}

func generator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int

	after := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Time out")
		case <-tick:
			fmt.Println("len(values) = ", len(values))
		case <-after:
			fmt.Println("Finished")
			os.Exit(1)
		}
	}
}
