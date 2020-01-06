package zhenai

import (
	"crawler/engine"
	"crawler/parser/zhenai"
	"crawler_distributed/config"
)

const begin = "http://www.zhenai.com/zhenghun"

func Seed() []engine.Request {
	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:     begin,
		Handler: engine.NewNormalParserFunc(zhenai.CityListParser, config.CityListParser),
	})
	return requests
}
