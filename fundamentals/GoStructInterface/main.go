package main

import "fmt"

type show interface {
	toShow() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

func (p person) toShow() string {
	return p.name + " " + p.surname
}

func (p product) toShow() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func showTo(x show) {
	fmt.Println(x.toShow())
}

func main() {
	var thing = person{
		name:    "Wene",
		surname: "Alves de Oliveira",
	}
	fmt.Println(thing.toShow())
	showTo(thing)

	thing2 := product{"Notebook", 2000.0}

	fmt.Println(thing2.toShow())
	showTo(thing2)
}
