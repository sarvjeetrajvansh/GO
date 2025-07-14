package main

import "fmt"

func arraySlice() {
	var arr [3]int // array of 3 integers, all zero-initialized
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	var processNumber = [10]int{10, 11, 12, 13, 1, 2, 3, 14, 15, 7}
	letters := [5]string{"S", "h", "i", "v", "a"}
	auto := [...]int{11, 21, 31} // length is 3
	fmt.Println()
	for _, val := range processNumber {
		fmt.Print(val, ",")
	}
	fmt.Println()
	for _, val := range auto {
		fmt.Print(val, ",")
	}
	fmt.Println()
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], ",")
	}
	fmt.Println()
	fmt.Println("Letters:", letters)

	// Slices
	s := []int{10, 20, 30}
	fmt.Println("Slice before append:", s)
	s = append(s, 40)
	fmt.Println("Slice after append:", s)

	// Slicing arrays
	a := [5]int{100, 200, 300, 400, 500}
	slice := a[1:4] // 200, 300, 400
	fmt.Println("Slice of array:", slice)
	slice[0] = 999
	fmt.Println("Modified array:", a)

	// Using make
	d := make([]int, 2, 5)
	d[0] = 42
	d[1] = 43
	d = append(d, 44)
	fmt.Println("Slice with make:", d)
}
