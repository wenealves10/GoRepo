package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getNameComplete() {
	fmt.Println(p.name + " " + p.surname)
}

func (p *person) setNameComplete(name string) {
	parts := strings.Split(name, " ")
	p.name = parts[0]
	p.surname = strings.Join(parts[1:], " ")
}

func main() {
	personOne := person{name: "Wene", surname: "Alves"}
	personOne.getNameComplete()
	personOne.setNameComplete("Wene Alves de Oliveira")
	personOne.getNameComplete()
}
