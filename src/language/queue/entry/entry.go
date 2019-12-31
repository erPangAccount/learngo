package main

import (
	"fmt"
	queue2 "language/queue"
)

func main() {
	q := queue2.Queue{0}

	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())     //0
	fmt.Println(q.Pop())     //1
	fmt.Println(q.IsEmpty()) //false
	fmt.Println(q.Pop())     //2
	fmt.Println(q.IsEmpty()) //true

	q.Push("asdf")
	fmt.Println(q.Pop())
}
