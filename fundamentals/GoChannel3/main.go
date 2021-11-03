package main

import (
	"fmt"
	"time"
)

func rotine(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	fmt.Println("Execute....")
	channel <- 5
	channel <- 6
	channel <- 7
}

func main() {
	ch := make(chan int, 3)

	go rotine(ch)

	time.Sleep(time.Second)

	fmt.Println(<-ch)

}
