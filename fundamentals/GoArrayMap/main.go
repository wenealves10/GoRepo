package main

import "fmt"

func main() {
	// var noteStudents map[int]string
	noteStudents := make(map[int]string)

	noteStudents[12211212112] = "Weene Alves"
	noteStudents[11225552252] = "Mateus Silvas"
	noteStudents[12254544522] = "Ismael Albert"

	fmt.Println(noteStudents)

	for cpf, name := range noteStudents {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Printf("%s Foi deletado!\n", noteStudents[12254544522])
	delete(noteStudents, 12254544522)
	for cpf, name := range noteStudents {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}
	// fmt.Println(noteStudents)
}
