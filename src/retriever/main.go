package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
)

type RetrieverInterface interface {
	Get(url string) string
}

func download(r RetrieverInterface) string {
	return r.Get("http://baidu.com")
}

func main() {
	var r RetrieverInterface	//声明变量接口
	r = mock.Retriever{"mockRetriever"} //把实现放入接口变量中
	fmt.Println(download(r))

	r = real2.RetrieverStruct{}
	fmt.Println(download(r))
}

