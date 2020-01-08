package main

import (
	"crawler/engine/zhenai"
	"crawler/scheduler"
	zhenai2 "crawler/seed/zhenai"
	client2 "crawler_distributed/itemRpc/client"
	rpc2 "crawler_distributed/rpc"
	"crawler_distributed/workerRpc/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemServiceHost    = flag.String("itemHost", "", "item service host")
	workerServiceHosts = flag.String("workerHosts", "", "worker service hosts")
)

func main() {
	flag.Parse()
	if *itemServiceHost == "" {
		panic("item service host can not empty!")
	}

	if *workerServiceHosts == "" {
		panic("worker service hosts can not empty!")
	}

	itemChan, err := client2.ItemServer(*itemServiceHost)
	if err != nil {
		panic(err)
	}

	workerClient := client.CreateWorkerClient(createWorkerClients(strings.Split(*workerServiceHosts, ",")))

	e := zhenai.QueueEngine{
		Scheduler:    &scheduler.QueueScheduler{},
		WorkerCount:  10,
		ItemChan:     itemChan,
		WorkerClient: workerClient,
	}

	e.Run(zhenai2.Seed())
}

func createWorkerClients(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		rpcClient, err := rpc2.NewRpcClient(host)
		if err != nil {
			log.Printf("create service: %s's client is fail; Error info: %v", host, err)
			continue
		}
		clients = append(clients, rpcClient)
	}

	rpcClientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				rpcClientChan <- c
			}
		}
	}()
	return rpcClientChan
}
