package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numberRandom(a int) int {
	number := rand.NewSource(time.Now().UnixNano())
	x := rand.New(number)

	return x.Intn(a)
}

func main() {

	number := 0

	for number <= 100000 {
		fmt.Println("Valor: ", number)
		number = numberRandom(100000)
		time.Sleep(time.Second)
	}

}
