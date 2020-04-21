package engine

type Scheduler interface {
	ReadyNotifier
	Register(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	ChanCount int
	ItemChan  chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.ChanCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Register(r)
	}

	for {
		result := <-out
		// 遍历打印item
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		//将Request注册
		for _, r := range result.Requests {
			if !isDuplicate(r.URL) {
				e.Scheduler.Register(r)
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
