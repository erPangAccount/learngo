package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "http://taobao.com"
	response, err := http.NewRequest(http.MethodGet, url, nil)
	response.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")
	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req)
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}

	resp, err := client.Do(response)
	//resp, err := http.DefaultClient.Do(response)

	//resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}
