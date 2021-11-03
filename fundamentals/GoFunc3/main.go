package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}
func sub(a, b int) int {
	return a - b
}

func exec(fun func(int, int) int, a, b int) int {
	return fun(a, b)
}

func main() {
	fmt.Printf("Soma %d + %d = %d\n", 4, 5, exec(sum, 4, 5))
	fmt.Printf("Subtração %d - %d = %d\n", 5, 4, exec(sub, 5, 4))
	fmt.Printf("Divisão %d ÷ %d = %d\n", 10, 5, exec(div, 10, 5))
	fmt.Printf("Multiplicação %d x %d = %d\n", 4, 5, exec(mult, 4, 5))
}
