package main

import "fmt"

type person struct {
	name string
}

type dynamic interface{}

func main() {
	var thing interface{}
	thing = 10
	fmt.Println(thing)
	thing = "Weneeee"
	fmt.Println(thing)
	thing = true
	fmt.Println(thing)
	var thing2 dynamic = "Eeee"
	fmt.Println(thing2)
	thing2 = person{"Wene ALves"}
	fmt.Println(thing2)
}
