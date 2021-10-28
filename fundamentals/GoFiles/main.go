package main

import (
	"fmt"

	"./math"
)

func main() {
	result := sum(2, 5)
	fmt.Printf("%T \n", result)
	fmt.Printf("%d \n", result)

	result = math.Sum(8, 7)
	fmt.Printf("%v \n", result)
	fmt.Printf("%.2f \n", math.PI)
	fmt.Printf("%d \n", math.SumX(70))
}

func sum(a int, b int) int {
	return a + b
}
