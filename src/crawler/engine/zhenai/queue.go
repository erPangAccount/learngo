package zhenai

import (
	"crawler/engine"
)

type QueueEngine struct {
	Scheduler    Scheduler
	WorkerCount  int
	ItemChan     chan engine.Item
	WorkerClient WorkerClient
}

var visitedUrls = make(map[string]string)

type WorkerClient func(engine.Request) (engine.RequestResult, error)

func (q *QueueEngine) Run(seeds []engine.Request) {
	if q.WorkerCount < 1 {
		panic("dont have worker work")
	}

	out := make(chan engine.RequestResult)
	q.Scheduler.Run()
	for i := 0; i < q.WorkerCount; i++ {
		q.createQueueWorker(q.Scheduler.ReturnWorkChan(), out, q.Scheduler)
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
			if !existsVisitedUrls(request.Url) {
				q.Scheduler.Submit(request)
			}
		}
	}
}

func (q *QueueEngine) createQueueWorker(in chan engine.Request, out chan engine.RequestResult, s ReadNotify) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := q.WorkerClient(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/**
判断是否已经获取过此网页相关信息了
*/
func existsVisitedUrls(url string) bool {
	if _, ok := visitedUrls[url]; ok {
		return true
	}
	visitedUrls[url] = url
	return false
}
