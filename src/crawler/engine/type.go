package engine

import (
	"fmt"
	"reflect"
	"runtime"
)

const ElasticHost = "http://192.168.0.103:9200/"

type Item struct {
	Url    string
	Type   string
	Id     string
	DoType interface{}
}

type ParserFunc func([]byte, string) RequestResult

type Request struct {
	Url     string
	Handler ParserFunc
}

type RequestResult struct {
	Requests []Request
	Items    []Item
}

func NilRequestResultFunc([]byte, string) RequestResult {
	return RequestResult{}
}

func GetFuncName(f interface{}) string {
	if typeName := reflect.TypeOf(f).Kind(); typeName != reflect.Func {
		panic(fmt.Sprintf("the f type must be func;but got %s", typeName))
	}

	ptr := reflect.ValueOf(f).Pointer()
	return runtime.FuncForPC(ptr).Name()
}
