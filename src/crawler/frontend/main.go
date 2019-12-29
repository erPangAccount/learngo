package main

import (
	"crawler/frontend/controller"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src/crawler/frontend/view")))

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		file, e := os.Open("./src/crawler/frontend/view/searchResult.html")
		if e != nil {
			panic(e)
		}

		bytes, e := ioutil.ReadAll(file)
		if e != nil {
			panic(e)
		}

		_, e = fmt.Fprint(writer, fmt.Sprintf("%s", bytes))
		if e != nil {
			panic(e)
		}
	})

	searchResultHandler, err := controller.CreateSearchResultHandler("./src/crawler/frontend/view/searchResult.html")
	if err != nil {
		panic(err)
	}
	http.Handle("/search", searchResultHandler)

	port := ":8888"
	serve := http.ListenAndServe(port, nil)
	if serve != nil {
		panic(serve)
	}
}
