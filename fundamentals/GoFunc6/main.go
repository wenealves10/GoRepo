package main

import (
	"fmt"
	"log"
)

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Number invalid %d", n)
	case n == 0:
		return 1, nil
	default:
		factorialAfter, _ := factorial(n - 1)
		return n * factorialAfter, nil
	}
}

func main() {
	number := 0

	fmt.Print("Digite um número: ")
	fmt.Scanln(&number)
	factorialNumber, err := factorial(number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(factorialNumber)
}
