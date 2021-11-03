package main

import "fmt"

type sports interface {
	onTurbo()
}

type car struct {
	name    string
	model   string
	year    int
	chassis string
}

type ferrari struct {
	car
	turbo bool
}

func (f *ferrari) onTurbo() {
	f.turbo = true
}

func main() {
	carOne := ferrari{
		car: car{
			name:    "F40",
			model:   "Ferrari XPTO",
			year:    2021,
			chassis: "XAKGGAYT24KKSA7",
		},
		turbo: false,
	}
	carOne.onTurbo()

	var cartwo sports = &ferrari{
		car: car{
			name:    "F42",
			model:   "Ferrari XPTO2",
			year:    2022,
			chassis: "XAKGGA112230KKSA7",
		},
		turbo: false,
	}

	cartwo.onTurbo()

	fmt.Println(carOne, cartwo)
}
