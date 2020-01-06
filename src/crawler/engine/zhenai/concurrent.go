package zhenai

import (
	"crawler/engine"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadNotify
	Submit(engine.Request)
	ReturnWorkChan() chan engine.Request
	Run()
}

type ReadNotify interface {
	WorkerReady(chan engine.Request)
}

func (c *ConcurrentEngine) Run(requests []engine.Request) {
	if c.WorkerCount < 1 {
		panic("dont has worker to do ！")
	}

	out := make(chan engine.RequestResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.ReturnWorkChan(), out)
	}

	for _, request := range requests {
		c.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item interface{}) { c.ItemChan <- item }(item)
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan engine.Request, out chan engine.RequestResult) {
	go func() {
		for {
			request := <-in
			result, e := Worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
