package zhenai

import (
	"crawler/engine"
	"crawler/parser/zhenai"
)

const begin = "http://www.zhenai.com/zhenghun"

func Seed() []engine.Request {
	var requests []engine.Request
	requests = append(requests, engine.Request{
		Url:     begin,
		Handler: zhenai.CityListParser,
	})
	return requests
}
