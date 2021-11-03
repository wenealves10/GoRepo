package main

import "fmt"

type note float64

func (n note) entries(initial, fim float64) bool {
	return float64(n) >= initial && float64(n) <= fim
}

func studentsOfNotes(n note) string {
	switch {
	case n.entries(9, 10):
		return "A"
	case n.entries(7, 8):
		return "B"
	case n.entries(5, 6):
		return "C"
	case n.entries(3, 4):
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(studentsOfNotes(10))
	fmt.Println(studentsOfNotes(7))
	fmt.Println(studentsOfNotes(2))
}
