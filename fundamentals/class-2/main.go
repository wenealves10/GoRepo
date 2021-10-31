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

	for {
		fmt.Print("Digite um número: ")
		fmt.Scanln(&userNumber)

		if number == userNumber {
			fmt.Print("Você acertou!!!\n")
			break
		} else if number < userNumber {
			fmt.Println("Você chutou acima!!!")
		} else if number > userNumber {
			fmt.Println("Você chutou abaixo!!!")
		}
	}

}
