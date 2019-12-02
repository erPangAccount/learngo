package main

import (
	"errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

/**
定义用户错误接口
*/
type userError interface {
	error
	Message() string
}

type errHandlerType func(writer http.ResponseWriter, request *http.Request) error

/**
函数式编程
*/
func errHandler(handlerType errHandlerType) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			recoverErr := recover()
			if recoverErr != nil {
				log.Println(recoverErr)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		err := handlerType(writer, request)
		if err != nil {
			log.Printf("Error request %v\n response: %s", request, err)
			code := http.StatusOK

			/**
			处理用户可见错误信息
			*/
			if userError, ok := err.(userError); ok { //type assertion
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			//处理用户不可见错误
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errHandler(filelisting.Handle))

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello Go Web"))
	//})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
