package zhenai

import (
	"crawler/engine"
	"crawler/fetche"
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seed []engine.Request) {
	var requests []engine.Request
	requests = append(requests, seed...)

	for len(requests) > 0 {
		targetRequest := requests[0]
		requests = requests[1:]

		//获取开始地址
		result, err := worker(targetRequest)
		if err != nil {
			continue
		}

		for i, request := range result.Requests {
			log.Printf("Url: %s; Handler: %v; Item: %s \n", request.Url, engine.GetFuncName(request.Handler), result.Items[i])
		}
		requests = append(requests, result.Requests...)
	}
}

func worker(request engine.Request) (engine.RequestResult, error) {
	contents, err := fetche.Fetcher(request.Url)
	if err != nil {
		log.Println(err)
		return engine.RequestResult{}, err
	}

	return request.Handler(contents, request.Url), nil
}
