package engine

import (
	"fmt"
	"reflect"
	"runtime"
)

type Request struct {
	Url     string
	Handler func([]byte) RequestResult
}

type RequestResult struct {
	Requests []Request
	Items    []interface{}
}

func NilRequestResultFunc([]byte) RequestResult {
	return RequestResult{}
}

func GetFuncName(f interface{}) string {
	if typeName := reflect.TypeOf(f).Kind(); typeName != reflect.Func {
		panic(fmt.Sprintf("the f type must be func;but got %s", typeName))
	}

	ptr := reflect.ValueOf(f).Pointer()
	return runtime.FuncForPC(ptr).Name()
}
