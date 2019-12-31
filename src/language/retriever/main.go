package main

import (
	"fmt"
	mock2 "language/retriever/mock"
	real3 "language/retriever/real"
	"time"
)

type RetrieverInterface interface {
	Get(url string) string
}

func download(r RetrieverInterface) string {
	return r.Get("http://baidu.com")
}

func inspect(r RetrieverInterface) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock2.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real3.RetrieverStruct:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r RetrieverInterface             //声明变量接口
	r = mock2.Retriever{"mockRetriever"} //把实现放入接口变量中
	//fmt.Println(download(r))
	inspect(r)

	r = &real3.RetrieverStruct{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)

	// type assertion
	r = &real3.RetrieverStruct{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//严格方式
	realRetriever := r.(*real3.RetrieverStruct)
	fmt.Println(realRetriever.TimeOut) //1m0s

	r = mock2.Retriever{"mockRetriever"} //把实现放入接口变量中
	//宽松方式
	if mockRetriever, ok := r.(mock2.Retriever); ok { //not is mockRetriever
		fmt.Println(mockRetriever)
	} else {
		fmt.Println("not is mockRetriever")
	}
}
