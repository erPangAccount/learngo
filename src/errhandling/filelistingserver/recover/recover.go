package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err) //this is an error
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))

	panic(123) //panic: 123 [recovered]
	//	panic: 123
}

func main() {
	tryRecover()
}
