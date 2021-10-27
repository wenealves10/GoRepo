package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 5.9
	notalF := int(nota)
	fmt.Println(notalF)

	//cuidados...
	// retorna codigo da tabela asc
	fmt.Println("Test " + string(rune((97))))

	//int to string
	fmt.Println("Test" + strconv.Itoa(123))

	//string to int

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string to bool
	// false, true, 0, 1
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdade...")
	}
}
