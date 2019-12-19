package zhenai

import (
	"crawler/engine"
	"log"
)

type QueueEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (q *QueueEngine) Run(seeds []engine.Request) {
	if q.WorkerCount < 1 {
		panic("dont have worker work")
	}
	q.Scheduler.Run()

	out := make(chan engine.RequestResult)
	for i := 0; i < q.WorkerCount; i++ {
		createQueueWorker(out, q.Scheduler)
	}

	for _, request := range seeds {
		q.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("%v", item)
		}

		for _, request := range result.Requests {
			q.Scheduler.Submit(request)
		}
	}
}

func createQueueWorker(out chan engine.RequestResult, s Scheduler) {
	in := make(chan engine.Request)
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
