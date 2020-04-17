package scheduler

import "github.com/corrots/go-demo/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Register(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}
