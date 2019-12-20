package scheduler

import (
	"crawler/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) ReturnWorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workChan = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request

		for {
			var targetRequestChan engine.Request
			var targetWorkerChan chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				targetRequestChan = requestQueue[0]
				targetWorkerChan = workerQueue[0]
			}

			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workChan:
				workerQueue = append(workerQueue, w)
			case targetWorkerChan <- targetRequestChan:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
