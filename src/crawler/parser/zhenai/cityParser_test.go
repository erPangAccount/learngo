package zhenai_test

import (
	"crawler/parser/zhenai"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCityParser(t *testing.T) {
	contents, err := ioutil.ReadFile("cityParserTemplate.html")
	if err != nil {
		panic(err)
	}

	result := zhenai.CityParser(contents)

	const resultSize = 25

	var expectedUrls = []string{
		"http://album.zhenai.com/u/1662184411",
		"http://album.zhenai.com/u/1486293757",
		"http://album.zhenai.com/u/1197897297",
	}

	var expectedItems = []string{ // 由于前5个是随机会变得，所以从第6个开始校验
		"时间",
		"流浪雪",
		"强哥",
	}

	// verify result
	if len(result.Requests) != resultSize {
		t.Errorf("result have %d; but got %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result have %d; but got %d", resultSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		tarI := i + 5
		if url != result.Requests[tarI].Url {
			t.Errorf("expected url #%d: %s; but got %s", tarI, url, result.Requests[tarI].Url)
		}
	}

	for i, city := range expectedItems {
		tarI := i + 5
		if city != fmt.Sprintf("%s", result.Items[tarI]) {
			t.Errorf("expected url #%d: %s; but got %s", tarI, city, result.Items[tarI])
		}
	}

}
