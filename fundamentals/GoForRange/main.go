package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 5, 8}

	for i, num := range numbers {
		fmt.Printf("index: %d  element: %d\n", i+1, num)
	}

	for _, num := range numbers {
		fmt.Printf(" %d", num)
	}

}
