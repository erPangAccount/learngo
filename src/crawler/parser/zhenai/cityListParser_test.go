package zhenai_test

import (
	"crawler/parser/zhenai"
	"io/ioutil"
	"testing"
)

func TestCityListParser(t *testing.T) {
	contents, err := ioutil.ReadFile("cityListParserTemplate.html")
	if err != nil {
		panic(err)
	}

	result := zhenai.CityListParser(contents, "")

	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	// verify result

	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expected url #%d: %s; but got %s", i, url, result.Requests[i].Url)
		}
	}

}
