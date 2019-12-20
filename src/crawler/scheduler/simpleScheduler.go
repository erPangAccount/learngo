package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) ReturnWorkChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workChan <- request }()
}
