package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)

	err = errors.New("this is custom err")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathErr.Op, pathErr.Path, pathErr.Err, pathErr.Error())
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("test.txt")
}
