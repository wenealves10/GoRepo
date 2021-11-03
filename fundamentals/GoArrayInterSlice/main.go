package main

import "fmt"

func main() {
	arrayOriginal := make([]int, 10, 20)
	arraySlice := append(arrayOriginal, 1, 2, 3, 4, 5, 6)
	fmt.Println(arrayOriginal, arraySlice)

	arrayOriginal[1] = 10

	fmt.Println(arrayOriginal, arraySlice)

}
