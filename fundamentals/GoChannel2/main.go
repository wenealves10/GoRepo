package main

import (
	"fmt"
	"time"
)

func mult(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base
	time.Sleep(time.Second)
	channel <- 4 * base
	time.Sleep(time.Second * 5)
	channel <- 12 * base
}

func main() {
	c := make(chan int)
	go mult(10, c)

	fmt.Println("Startup....")

	a, b := <-c, <-c

	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println("Fim...")
}
