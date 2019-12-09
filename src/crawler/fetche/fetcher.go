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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httpCode err: ", resp.StatusCode)
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
