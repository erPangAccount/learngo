package engine

import (
	"crawler_distributed/config"
	"fmt"
	"reflect"
	"runtime"
)

const ElasticHost = "http://192.168.12.13:9200/"

type Item struct {
	Url    string
	Type   string
	Id     string
	DoType interface{}
}

type ParserFunc func([]byte, string) RequestResult

type Parser interface {
	Parser(contents []byte, url string) RequestResult
	Serialize() (name string, ags interface{})
}

type Request struct {
	Url     string
	Handler Parser
}

type RequestResult struct {
	Requests []Request
	Items    []Item
}

type NilRequestResultFunc struct{}

func (n NilRequestResultFunc) Parser(_ []byte, _ string) RequestResult {
	return RequestResult{}
}

func (n NilRequestResultFunc) Serialize() (name string, ags interface{}) {
	return config.NilRequestResultFunc, nil
}

type NormalParserFunc struct {
	parser ParserFunc
	name   string
}

func (n *NormalParserFunc) Parser(contents []byte, url string) RequestResult {
	return n.parser(contents, url)
}

func (n *NormalParserFunc) Serialize() (name string, ags interface{}) {
	return n.name, nil
}

func NewNormalParserFunc(parser ParserFunc, name string) *NormalParserFunc {
	return &NormalParserFunc{
		parser: parser,
		name:   name,
	}
}

func GetFuncName(f interface{}) string {
	if typeName := reflect.TypeOf(f).Kind(); typeName != reflect.Func {
		panic(fmt.Sprintf("the f type must be func;but got %s", typeName))
	}

	ptr := reflect.ValueOf(f).Pointer()
	return runtime.FuncForPC(ptr).Name()
}
