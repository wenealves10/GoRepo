package main

import "fmt"

func listOfStudents(students ...string) {
	fmt.Println("----------------------------------------")
	fmt.Println("         Lista de Aprovados             ")
	fmt.Println("----------------------------------------")
	for i, student := range students {
		fmt.Printf("%d) %s\n", i+1, student)
	}
	fmt.Println("----------------------------------------")
}

func main() {
	students := []string{"Wene Alves", "Ismael Albert", "Eduardo Mota", "Amilson Silva", "Nadson Cruz"}
	listOfStudents(students...)
}
