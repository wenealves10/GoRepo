package main

import (
	"log"
	"math/rand"
	"time"
)

func numberRandom(a int) int {
	number := rand.NewSource(time.Now().UnixNano())
	x := rand.New(number)

	return x.Intn(a)
}

func main() {
	if i := numberRandom(10); i < 5 {
		log.Println("Você Ganhou!!, Valor: ", i)
	} else {
		log.Println("Você Perdeu!!, Valor: ", i)
	}
}
