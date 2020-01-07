package main

import (
	"crawler/engine"
	"crawler/engine/zhenai"
	"crawler/persist"
	"crawler/scheduler"
	zhenai2 "crawler/seed/zhenai"
)

func main() {
	//zhenai.SimpleEngine{}.Run(zhenai2.Seed())
	//e := zhenai.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//	ItemChan:    persist.ItemServer(),
	//}
	itemChan, err := persist.ItemServer("test")
	if err != nil {
		panic(err)
	}

	e := zhenai.QueueEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		WorkerClient: func(request engine.Request) (result engine.RequestResult, e error) {
			return zhenai.Worker(request)
		},
	}

	e.Run(zhenai2.Seed())
}
