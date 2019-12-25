package zhenai

import (
	"crawler/engine"
)

type QueueEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan engine.Item
}

func (q *QueueEngine) Run(seeds []engine.Request) {
	if q.WorkerCount < 1 {
		panic("dont have worker work")
	}

	out := make(chan engine.RequestResult)
	q.Scheduler.Run()
	for i := 0; i < q.WorkerCount; i++ {
		createQueueWorker(q.Scheduler.ReturnWorkChan(), out, q.Scheduler)
	}

	for _, request := range seeds {
		q.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item engine.Item) {
				q.ItemChan <- item
			}(item)
		}

		for _, request := range result.Requests {
			q.Scheduler.Submit(request)
		}
	}
}

func createQueueWorker(in chan engine.Request, out chan engine.RequestResult, s ReadNotify) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
