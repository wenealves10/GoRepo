package main

import "fmt"

func main() {
	students := map[string]float64{
		"Wene Alves":    2000,
		"Ismael Albert": 2000,
		"Eduardo Mota":  2000,
	}

	students["Marcelo Silva"] = 5000

	delete(students, "Invalid")

	for name, salt := range students {
		fmt.Printf("Nome: %s\nSalário: %.2f\n", name, salt)
		fmt.Println("-------------------------------------")
	}
}
