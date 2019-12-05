package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
)

const begin = "http://www.zhenai.com/zhenghun"

type Request struct {
	Url     string
	Handler func(Request) RequestResult
}

type RequestResult struct {
	Requests []Request
	Items    []interface{}
}

func nilRequestResultFunc(request Request) RequestResult {
	return RequestResult{
		Requests: nil,
		Items:    nil,
	}
}

var Q []Request

func main() {
	//获取开始地址
	resp, err := http.Get(begin)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("httpCode err: ", resp.StatusCode)
		return
	}

	respBodyReader := bufio.NewReader(resp.Body)
	bytes, err := respBodyReader.Peek(1024)
	if err != nil {
		panic(err)
	}
	bodyEncode := charsetToUtf8(bytes)

	ut8Reader := transform.NewReader(respBodyReader, bodyEncode.NewDecoder())
	contents, err := ioutil.ReadAll(ut8Reader)
	if err != nil {
		panic(err)
	}

	result := getCityAndUrl(contents)
	for i, request := range result.Requests {
		fmt.Printf("Url: %s; Handler: %v; City: %s \n", request.Url, getFuncName(request.Handler), result.Items[i])
	}

	fmt.Println("------------Queue------------")
	for _, item := range Q {
		fmt.Println(item)
	}
	fmt.Println(len(Q))
}

func charsetToUtf8(b []byte) encoding.Encoding {
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}

var cityRe *regexp.Regexp = regexp.MustCompile(`<a [_a-z\-\=\" ]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)

func getCityAndUrl(contents []byte) RequestResult {
	var requestResult RequestResult
	citySlice := cityRe.FindAllSubmatch(contents, -1)
	for _, val := range citySlice {
		request := Request{
			Url:     string(val[1]),
			Handler: nilRequestResultFunc,
		}
		requestResult.Requests = append(requestResult.Requests, request)
		Q = append(Q, request)
		requestResult.Items = append(requestResult.Items, val[2])
	}

	return requestResult
}

func getFuncName(f interface{}) string {
	if typeName := reflect.TypeOf(f).Kind(); typeName != reflect.Func {
		panic(fmt.Sprintf("the f type must be func;but got %s", typeName))
	}

	ptr := reflect.ValueOf(f).Pointer()
	return runtime.FuncForPC(ptr).Name()
}
