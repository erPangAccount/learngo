package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 011
var bb = "asdf"

var (
	a = 1
	b = "aaa"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //定义后都有一个初始值
}

func variableInitValue() {
	var a int = 1
	var s string = "aaaa"
	fmt.Println(a, s)
}

func variableTypeAuto() {
	var a, b, s = 1, true, "aaa"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, s := 1, true, "short"
	fmt.Println(a, b, s)
}

func caclcTringle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

/**
欧拉公式
*/
func euler() {
	value := cmplx.Exp(1i*math.Pi) + 1 // Exp表示e的多少次方
	fmt.Println(value)
	fmt.Printf("%.3f", value)
}

/**
强制类型转换
*/
func forcedTypeConversion() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
