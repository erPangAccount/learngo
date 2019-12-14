package zhenai_test

import (
	"crawler/parser/zhenai"
	"fmt"
	"io/ioutil"
	"testing"
)

const resultSize = 494

var expectedUrls = []string{
	"http://www.zhenai.com/zhenghun/aba",
	"http://www.zhenai.com/zhenghun/akesu",
	"http://www.zhenai.com/zhenghun/alashanmeng",
}
var expectedCities = []string{
	"阿坝",
	"阿克苏",
	"阿拉善盟",
}

func TestCityListParser(t *testing.T) {
	contents, err := ioutil.ReadFile("cityListParserTemplate.html")
	if err != nil {
		panic(err)
	}

	result := zhenai.CityListParser(contents)

	// verify result
	if len(result.Requests) != resultSize {
		t.Errorf("result have %d; but got %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result have %d; but got %d", resultSize, len(result.Items))
	}

	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expected url #%d: %s; but got %s", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCities {
		if city != fmt.Sprintf("%s", result.Items[i]) {
			t.Errorf("expected url #%d: %s; but got %s", i, city, result.Items[i])
		}
	}

}
