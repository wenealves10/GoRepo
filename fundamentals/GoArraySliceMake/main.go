package main

import "fmt"

func main() {
	array := make([]int, 10)

	array[2] = 10

	fmt.Println(array)

	array = make([]int, 10, 20)

	fmt.Println(array, len(array), cap(array))

	array = append(array, 1, 2, 3, 3, 2, 5, 5, 56, 6, 6, 87)

	fmt.Println(array, len(array), cap(array))

}
