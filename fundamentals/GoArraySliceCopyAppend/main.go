package main

import "fmt"

func main() {
	array := [3]int{100, 23, 3}
	var slice []int

	// array = append(array, 1, 2, 3, 3)

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(array, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)

	fmt.Println(slice2, slice)

}
