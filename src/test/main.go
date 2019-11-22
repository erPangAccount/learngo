package main

import "fmt"

type A struct {
	a int
}
type B struct {
	a int
}
type C struct {
	A
	B
}

func main() {
	c := C{}
	c.A.a = 1
	c.B.a = 1
	fmt.Println(c)
}
