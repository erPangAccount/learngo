package main

import (
	"crawler_distributed/rpc"
	"crawler_distributed/workerRpc"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "service port")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("must have a service port!")
	}

	log.Fatal(rpc.ServeRpc(workerRpc.WorkService{}, fmt.Sprintf(":%d", *port)))
}
