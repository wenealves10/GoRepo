package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Sleeping for %v\n", dur)
		time.Sleep(dur)
	}
	fmt.Println("Done!")
}
