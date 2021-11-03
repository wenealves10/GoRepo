package main

import "fmt"

func media(notes ...float64) float64 {
	total := 0.0

	if len(notes) == 0 {
		return 0
	}

	for _, note := range notes {
		total += note
	}

	return total / float64(len(notes))
}

func main() {
	fmt.Printf("Média: %.2f\n", media(7, 8, 9, 10))
	fmt.Println("Média: ", media())
}
