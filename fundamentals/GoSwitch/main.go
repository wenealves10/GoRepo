package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

func studyClass(n int) string {
	switch n {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	n := 0
	fmt.Print("Digite sua nota: ")
	fmt.Scanln(&n)
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println("Você está com nota: ", studyClass(n))
}
