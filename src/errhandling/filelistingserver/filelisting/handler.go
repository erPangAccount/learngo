package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/**
实现用户错误接口
*/
type userError string

func (userErr userError) Error() string {
	return userErr.Message()
}

func (userErr userError) Message() string {
	return string(userErr)
}

const requestFix = "/list/"

func Handle(writer http.ResponseWriter, request *http.Request) error {
	//判断外部传入是否为/list/开头的请求地址
	if strings.Index(request.URL.Path, requestFix) != 0 {
		return userError(fmt.Sprintf("the request prefix must be %s", requestFix))
	}

	path := request.URL.Path[len(requestFix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = writer.Write(contents)
	if err != nil {
		return err
	}

	return nil
}
