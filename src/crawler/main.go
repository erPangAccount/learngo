package main

import (
	"crawler/engine/zhenai"
	"crawler/scheduler"
	zhenai2 "crawler/seed/zhenai"
)

func main() {
	//zhenai.SimpleEngine{}.Run(zhenai2.Seed())
	//e := zhenai.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	e := zhenai.QueueEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
	}

	e.Run(zhenai2.Seed())
}
