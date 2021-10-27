package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func print(value int) {
	fmt.Print(value)
}

func main() {
	result := sum(4, 5)
	print(result)
}
