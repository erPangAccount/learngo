package zhenai

import (
	"crawler/engine"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
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

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Item: #%d: %s", itemCount, item)
			itemCount++
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
			result, e := worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
