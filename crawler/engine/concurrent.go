package engine

import "fmt"

type Scheduler interface {
	Register(request Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	ChanCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.ChanCount; i++ {
		e.createWorker(out)
	}

	for _, r := range seeds {
		e.Scheduler.Register(r)
	}

	itemCount := 0
	for {
		result := <-out
		// 遍历打印item
		for _, item := range result.Items {
			fmt.Printf("got item #%d: %v\n", itemCount, item)
			fmt.Printf("got item %v\n", item)
		}
		//将Request注册
		for _, r := range result.Requests {
			e.Scheduler.Register(r)
			itemCount++
		}
	}
}

func (e *ConcurrentEngine) createWorker(out chan ParseResult) {
	in := make(chan Request)
	go func() {
		for {
			e.Scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
