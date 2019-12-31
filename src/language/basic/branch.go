package main

import (
	"fmt"
)

func operation(operator string, a int, b int) int {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic(fmt.Sprintf("Error operator:%s", operator))
	}
	return result
}

func main() {
	fmt.Println(operation("+", 2, 2))
}
