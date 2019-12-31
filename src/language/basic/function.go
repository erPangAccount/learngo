package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	var result int
	var err error = nil
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result, _ = div(a, b)
	default:
		err = fmt.Errorf("Error operatioon: %s", op)
	}
	return result, err
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Callint function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	var s int
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(result)
	}

	q, r := div(13, 4)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 2, 3))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 3))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
