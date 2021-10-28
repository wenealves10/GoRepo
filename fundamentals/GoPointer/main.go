package main

import "fmt"

type Car struct {
	name  string
	year  int
	model string
}

func (c *Car) walk() {
	c.name = "Fiat uno 111"
	fmt.Println(c.name, "Walking...")
}

func main() {
	//memory-address-value <- a <- 10
	//memory-0xc00001a140-10 <- a <- 10
	a := 10

	fmt.Println(&a)

	var pointer *int = &a
	fmt.Println(pointer, *pointer)
	*pointer = 1000

	b := &a

	fmt.Println(a, *b, *pointer)

	suma(b)

	fmt.Println(a, *b, *pointer)

	car := Car{
		name:  "Fiat uno",
		year:  2001,
		model: "Gol",
	}

	car.walk()
	fmt.Print(car.name)
}

func suma(a *int) {
	*a = 200 + 1000
}
