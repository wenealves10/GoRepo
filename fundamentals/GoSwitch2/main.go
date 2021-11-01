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

func switchNumber(b int) string {
	switch {
	case b <= 10 && b >= 9:
		return "A"
	case b <= 8 && b >= 7:
		return "B"
	case b <= 6 && b >= 1:
		return "C"
	default:
		return "F"
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		fmt.Println(switchNumber(numberRandom(10)))
	}
}
