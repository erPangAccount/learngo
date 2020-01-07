package main

import (
	"crawler/engine/zhenai"
	"crawler/scheduler"
	zhenai2 "crawler/seed/zhenai"
	"crawler_distributed/config"
	client2 "crawler_distributed/itemRpc/client"
	"crawler_distributed/workerRpc/client"
)

func main() {
	itemChan, err := client2.ItemServer(config.ItemServiceHost)
	if err != nil {
		panic(err)
	}

	workerClient, err := client.CreateWorkerClient(config.WorkerServiceHost)
	if err != nil {
		panic(err)
	}

	e := zhenai.QueueEngine{
		Scheduler:    &scheduler.QueueScheduler{},
		WorkerCount:  10,
		ItemChan:     itemChan,
		WorkerClient: workerClient,
	}

	e.Run(zhenai2.Seed())
}
