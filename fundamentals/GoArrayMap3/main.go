package main

import "fmt"

func main() {
	students := map[string]map[string]float64{
		"A": {
			"Albert Nunes":  2000,
			"Arnoldo Silva": 5000,
		},
		"W": {
			"Wene Alves":   10000,
			"Wesley Silva": 1100,
		},

		"F": {
			"Frank": 2000,
		},
	}

	delete(students, "F")

	for letter, maps := range students {
		fmt.Println("------------------------------------------")
		fmt.Printf("Todos de Letra: %s\n", letter)
		for name, salt := range maps {
			fmt.Printf("Nome: %s Salário: %.2f\n", name, salt)
		}
		fmt.Println("------------------------------------------")
	}

}
