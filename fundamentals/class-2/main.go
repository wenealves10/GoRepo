package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)

	number := y1.Intn(10)
	var userNumber int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&userNumber)

	if number == userNumber {
		fmt.Print("Você acertou!!!\n")
	} else if number < userNumber {
		fmt.Printf("Você chutou acima!!! \nValor era: %d\n", number)
	} else if number > userNumber {
		fmt.Printf("Você chutou abaixo!!! \nValor era: %d\n", number)
	}

}
