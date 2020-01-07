package workerRpc

import (
	"crawler/engine"
	"crawler/model"
	"crawler/parser/zhenai"
	"crawler_distributed/config"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Handle SerializedParser
}

type RequestResult struct {
	Requests []Request
	Items    []engine.Item
}

//序列化
func SerializeRequest(request engine.Request) Request {
	name, ags := request.Handler.Serialize()
	return Request{
		Url: request.Url,
		Handle: SerializedParser{
			Name: name,
			Args: ags,
		},
	}
}

func SerializeRequestResult(r engine.RequestResult) RequestResult {
	result := RequestResult{
		Items: r.Items,
	}

	for _, request := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(request))
	}

	return result
}

//反序列化
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, e := DeserializeParser(r.Handle)
	if e != nil {
		return engine.Request{}, e
	}

	return engine.Request{
		Url:     r.Url,
		Handler: parser,
	}, nil
}

func DeserializeRequestResult(r RequestResult) engine.RequestResult {
	result := engine.RequestResult{}

	for _, item := range r.Items {
		doType, e := model.FromJsonObj(item.DoType)
		if e != nil {
			log.Printf("json decode error: %v", item.DoType)
			continue
		}
		item.DoType = doType
		result.Items = append(result.Items, item)
	}

	for _, req := range r.Requests {
		enginRequest, e := DeserializeRequest(req)
		if e != nil {
			log.Printf("deserializeRequestResult Error: %v", e)
			continue
		}

		result.Requests = append(result.Requests, enginRequest)
	}
	log.Println(result.Items)
	return result
}

func DeserializeParser(parser SerializedParser) (engine.Parser, error) {
	userName := ""
	if parser.Args != nil {
		name, ok := parser.Args.(string)
		if !ok {
			return nil, errors.Errorf("args err: %v", parser.Args)
		}
		userName = name
	}

	var parserFuncMap = map[string]engine.Parser{
		config.NilRequestResultFunc: engine.NilRequestResultFunc{},
		config.CityParser:           engine.NewNormalParserFunc(zhenai.CityParser, config.CityParser),
		config.UserInfoParser:       zhenai.NewUserInfoParser(userName),
		config.CityListParser:       engine.NewNormalParserFunc(zhenai.CityListParser, config.CityParser),
	}

	val, ok := parserFuncMap[parser.Name]
	if !ok {
		return nil, errors.Errorf("no parser: %v", parser.Name)
	}

	return val, nil
}
