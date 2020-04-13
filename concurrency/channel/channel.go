package done

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for v := range c {
		fmt.Printf("worker %d received %c\n", id, v)
	}
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		//channels[i] <- 'A' + i
	}

	for i := 0; i < 10; i++ {
		//channels[i] <- 'a' + i
		channels[i] <- 'A' + i
	}
}

func bufferedChan() {
	c := make(chan int, 1)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond * 10)
}

func closedChan() {
	c := make(chan int)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
}

func main() {
	chanDemo()
	//bufferedChan()
	//closedChan()
	time.Sleep(time.Millisecond)
}
