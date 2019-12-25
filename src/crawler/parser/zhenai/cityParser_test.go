package zhenai_test

import (
	"crawler/parser/zhenai"
	"io/ioutil"
	"testing"
)

func TestCityParser(t *testing.T) {
	contents, err := ioutil.ReadFile("cityParserTemplate.html")
	if err != nil {
		panic(err)
	}

	result := zhenai.CityParser(contents)

	var expectedUrls = []string{
		"http://album.zhenai.com/u/1662184411",
		"http://album.zhenai.com/u/1486293757",
		"http://album.zhenai.com/u/1197897297",
	}

	// verify result

	for i, url := range expectedUrls {
		tarI := i + 5
		if url != result.Requests[tarI].Url {
			t.Errorf("expected url #%d: %s; but got %s", tarI, url, result.Requests[tarI].Url)
		}
	}

}
