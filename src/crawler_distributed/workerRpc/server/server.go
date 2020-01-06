package main

import (
	"crawler_distributed/config"
	"crawler_distributed/rpc"
	"crawler_distributed/workerRpc"
	"log"
)

func main() {
	log.Fatal(rpc.ServeRpc(workerRpc.WorkService{}, config.WorkerServiceHost))
}
