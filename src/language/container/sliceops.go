package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%v, cap=%v\n", s, len(s), cap(s))
}

func main() {
	//create
	var s []int
	printSlice(s) //[], len=0, cap=0
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
	}
	printSlice(s) //[1 3 5 7 9 11 13 15 17 19], len=10, cap=16

	s1 := []int{2, 4, 6, 8}
	printSlice(s1) //[2 4 6 8], len=4, cap=4

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s3) //[0 0 0 0 0 0 0 0 0 0], len=10, cap=32

	//copy
	copy(s2, s1)
	printSlice(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16

	//delete
	//删除 8
	s2 = append(s2[:3], s[4:]...)
	printSlice(s2) //[2 4 6 9 11 13 15 17 19], len=9, cap=16
	//删除头部
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front) //2
	printSlice(s2)     //[4 6 9 11 13 15 17 19], len=8, cap=15

	s2Length := len(s2) - 1
	tail := s2[s2Length]
	s2 = s2[:s2Length]
	fmt.Println(tail) //19
	printSlice(s2)    //[4 6 9 11 13 15 17], len=7, cap=15
}
