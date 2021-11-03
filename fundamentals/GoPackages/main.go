package main

import "fmt"

func main() {
	p1 := Ponteir{2, 4}
	p2 := Ponteir{2, 10}

	fmt.Println(peccaries(p1, p2))
	fmt.Println(Distance(p1, p2))
}
