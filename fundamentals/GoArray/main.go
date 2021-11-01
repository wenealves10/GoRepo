package main

import "fmt"

func main() {
	var notas [4]float64
	var total float64

	for i := 0; i < len(notas); i++ {
		fmt.Printf("Digite a nota %d: ", i+1)
		fmt.Scanln(&notas[i])
	}

	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Sua média foi: %.2f\n", media)
}
