package client

import (
	"crawler/engine"
	"crawler/engine/zhenai"
	"crawler_distributed/config"
	"crawler_distributed/workerRpc"
	"net/rpc"
)

func CreateWorkerClient(clientChan chan *rpc.Client) zhenai.WorkerClient {
	// 调用服务
	return func(request engine.Request) (result engine.RequestResult, e error) {
		//序列化后传递给服务
		serializeRequest := workerRpc.SerializeRequest(request)
		var serializeResult workerRpc.RequestResult
		workerClient := <-clientChan
		err := workerClient.Call(config.WorkerService, serializeRequest, &serializeResult)
		if err != nil {
			return engine.RequestResult{}, nil
		}
		//将结果反序列化后返回
		return workerRpc.DeserializeRequestResult(serializeResult), nil
	}
}
