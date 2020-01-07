package main

import (
	"crawler_distributed/config"
	"crawler_distributed/rpc"
	"crawler_distributed/workerRpc"
	"log"
	"testing"
	"time"
)

func TestWorkService(t *testing.T) {
	const host = ":1235"
	//启服务
	go rpc.ServeRpc(workerRpc.WorkService{}, host)
	time.Sleep(time.Second)

	//启client
	client, err := rpc.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	//client调用
	var result workerRpc.RequestResult
	err = client.Call(config.WorkerService, workerRpc.Request{
		Url: "http://album.zhenai.com/u/86218455",
		Handle: workerRpc.SerializedParser{
			Name: config.UserInfoParser,
			Args: nil,
		},
	}, &result)
	if err != nil {
		panic(err)
	} else {
		requestResult := workerRpc.DeserializeRequestResult(result)
		log.Printf("%v", requestResult)
	}
}
