package main

import (
	"fmt"
)

func modify_arr(array [5]int) {
	array[0] = 10
	fmt.Println("in modify_arr, arrary value is:", array)
}

func print_info(my_slice []int) {
	fmt.Println("len:", len(my_slice))
	fmt.Println("cap:", cap(my_slice))
	for i, v := range my_slice {
		fmt.Println("element[", i, "]=", v)
	}
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s

}

func sliceInfo(x []int) {
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(x), cap(x), x)

}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify_arr(array)
	fmt.Println("in main, array value is :", array)
	my_slice1 := []int{1, 2, 3, 4, 5}
	my_slice2 := make([]int, 5)
	my_slice3 := make([]int, 5, 6)
	my_slice4 := append(my_slice3, 8, 9, 10)
	print_info(my_slice1)
	print_info(my_slice2)
	print_info(my_slice3)
	print_info(my_slice4)

	a := []int{1, 2, 3, 4, 5}
	var b = a[0:3]
	var c = [...]int{3, 6, 9, 2, 6, 4}
	d := c[0:2]

	sliceInfo(a)

	b[1] = 10
	b = append(b, 30)

	sliceInfo(a)
	sliceInfo(b)

	fmt.Printf("sum of b is %d\n", sum(b))
	fmt.Printf("sum of d is %d\n", sum(d))

}
