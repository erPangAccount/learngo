package main

import "fmt"

func updateSlice(a []int) {
	a[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] = ", arr[2:6]) //arr[2:6] =  [2 3 4 5]
	fmt.Println("arr[:6] = ", arr[:6])   //arr[:6] =  [0 1 2 3 4 5]
	fmt.Println("arr[2:] = ", arr[2:])   //arr[2:] =  [2 3 4 5 6 7 8]
	fmt.Println("arr[:] = ", arr[:])     //arr[:] =  [0 1 2 3 4 5 6 7 8]

	s1, s2 := arr[2:], arr[:]
	fmt.Println("s1=", s1, "s2=", s2, "arr=", arr) //s1= [2 3 4 5 6 7 8] s2= [0 1 2 3 4 5 6 7 8] arr= [0 1 2 3 4 5 6 7 8]
	updateSlice(s1)
	fmt.Println("s1=", s1, "s2=", s2, "arr=", arr) //s1= [100 3 4 5 6 7 8] s2= [0 1 100 3 4 5 6 7 8] arr= [0 1 100 3 4 5 6 7 8]

	s3 := arr[:]
	fmt.Println(s3) //[0 1 100 3 4 5 6 7 8]
	s3 = s3[:5]
	fmt.Println(s3) //[0 1 100 3 4]
	s3 = s3[2:]
	fmt.Println(s3) //[100 3 4]

	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2)) //s2=[5 6], len(s2)=2, cap(s2)=3
	//fmt.Println(s1[3:7]) // panic: runtime error: slice bounds out of range [:7] with capacity 6

	s1 = arr[2:6]
	s2 = s1[3:5]
	s3 = append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s1, s2, s3, s4, s5, arr)                                //[2 3 4 5] [5 6] [5 6 10] [5 6 10 11] [5 6 10 11 12] [0 1 2 3 4 5 6 10]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2)) //s2=[5 6], len(s2)=2, cap(s2)=3
	fmt.Printf("s3=%v, len(s3)=%v, cap(s3)=%v\n", s3, len(s3), cap(s3)) //s3=[5 6 10], len(s3)=3, cap(s3)=3
	fmt.Printf("s4=%v, len(s4)=%v, cap(s4)=%v\n", s4, len(s4), cap(s4)) //s4=[5 6 10 11], len(s4)=4, cap(s4)=6
	fmt.Printf("s5=%v, len(s5)=%v, cap(s5)=%v\n", s5, len(s5), cap(s5)) //s5=[5 6 10 11 12], len(s5)=5, cap(s5)=6
}
