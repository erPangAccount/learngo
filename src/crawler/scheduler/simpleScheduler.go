package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ReturnWorkChan(r chan engine.Request) {
	s.workChan = r
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workChan <- request }()
}
