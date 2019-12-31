package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fib() fibFunc {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibFunc func() int

func (fib fibFunc) Read(p []byte) (n int, err error) {
	value := fib()
	if value > 100000 {
		return 0, io.EOF
	}
	valueStr := fmt.Sprintf("%d ", value)
	return strings.NewReader(valueStr).Read(p)
}

func readContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f()) //1 1 2 3 5 8 13 21 34 55
	}
	fmt.Println()

	readContent(f)
}
