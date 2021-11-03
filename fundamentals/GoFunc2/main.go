package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sum = func(a, b int) int {
	return a + b
}

var numberRandom = func(n int) int {
	number := rand.NewSource(time.Now().UnixNano())
	newNumber := rand.New(number)

	return newNumber.Intn(n)
}

func main() {
	fmt.Println(sum(1, 2))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 2))

	for i := 0; i < 1000; i++ {

		numberOne := numberRandom(100)
		numberTwo := numberRandom(100)
		fmt.Println()
		fmt.Println("-----------------------------------------")
		fmt.Printf("%d + %d = %d\n", numberOne, numberTwo, sum(numberOne, numberTwo))
		fmt.Println("-----------------------------------------")
		fmt.Println()
	}
	for i := 0; i < 1000; i++ {
		numberOne := numberRandom(100)
		numberTwo := numberRandom(100)
		fmt.Println()
		fmt.Println("-----------------------------------------")
		fmt.Printf("%d - %d = %d\n", numberOne, numberTwo, sub(numberOne, numberTwo))
		fmt.Println("-----------------------------------------")
		fmt.Println()
	}
}
