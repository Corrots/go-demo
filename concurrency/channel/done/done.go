package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	//wg   *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
	}

	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}
	wg.Wait()

}

func main() {
	chanDemo()
	//time.Sleep(time.Millisecond)
}
