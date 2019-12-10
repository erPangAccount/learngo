package zhenai

import (
	"crawler/engine"
	"crawler/fetche"
	"fmt"
	"log"
)

func Run(seed []engine.Request) {
	var requests []engine.Request
	requests = append(requests, seed...)

	for len(requests) > 0 {
		targetRequest := requests[0]
		requests = requests[1:]

		//获取开始地址
		contents, err := fetche.Fetcher(targetRequest.Url)
		if err != nil {
			log.Println(err)
		}

		result := targetRequest.Handler(contents)
		for i, request := range result.Requests {
			fmt.Printf("Url: %s; Handler: %v; Item: %s \n", request.Url, engine.GetFuncName(request.Handler), result.Items[i])
		}
		requests = append(requests, result.Requests...)
	}
}
