package main

import "fmt"

func listOf(n int) int {
	defer fmt.Println("-------------------------")
	if n > 5000 {
		fmt.Println("List of Number Accepted")
		return n
	}

	fmt.Println("List of Number Rejected")
	return n
}

func main() {
	fmt.Println(listOf(10000))
	fmt.Println(listOf(4000))
}
