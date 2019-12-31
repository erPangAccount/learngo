package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println("----------test reflect struct----------------")
	student := Student{
		Name:  "Test",
		Age:   5,
		Score: 10,
	}
	structTest(&student)

	structTest(student)
}

// stuct refelct 测试

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Print() {
	fmt.Println(s)
}

func structTest(a interface{}) {
	val := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	kd := val.Kind()
	fmt.Println(val, t, kd)

	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		panic("error")
	}

}
