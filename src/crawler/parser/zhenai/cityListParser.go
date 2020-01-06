package zhenai

import (
	"crawler/engine"
	"crawler_distributed/config"
	"regexp"
)

var cityRe *regexp.Regexp = regexp.MustCompile(`<a [_a-z\-\=\" ]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)

func CityListParser(contents []byte, _ string) engine.RequestResult {
	var requestResult engine.RequestResult

	citySlice := cityRe.FindAllSubmatch(contents, -1)
	limit := 5
	for _, val := range citySlice {
		request := engine.Request{
			Url:     string(val[1]),
			Handler: engine.NewNormalParserFunc(CityParser, config.CityParser),
		}
		requestResult.Requests = append(requestResult.Requests, request)
		//requestResult.Items = append(requestResult.Items, val[2])
		limit--
		if limit == 0 {
			break
		}
	}
	return requestResult
}
