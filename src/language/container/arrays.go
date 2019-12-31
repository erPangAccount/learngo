package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var array1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var grid [4][5]int

	fmt.Println(array1, arr2, arr3)
	fmt.Println(grid)

	printArray(array1)
	fmt.Println(array1[0])
}
