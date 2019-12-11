package fetche

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	log.Printf("Fetching url: %s\n", url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new Request err: %s", err)
	}
	//添加user-agent信息,不添加此信息，获取珍爱网用户信息页面，会报403
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	//resp, err := http.Get(url)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httpCode err: %v", resp.StatusCode)
	}

	respBodyReader := bufio.NewReader(resp.Body)
	bytes, err := respBodyReader.Peek(1024)
	if err != nil {
		return nil, err
	}
	bodyEncode := charsetToUtf8(bytes)

	ut8Reader := transform.NewReader(respBodyReader, bodyEncode.NewDecoder())
	contents, err := ioutil.ReadAll(ut8Reader)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func charsetToUtf8(b []byte) encoding.Encoding {
	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
