package main

import (
	"fmt"
	"time"
)

func say(s string, done chan string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
	done <- "terminei"
}
func main() {
	done := make(chan string)
	go say("Hello World go lang", done)
	fmt.Println(<-done)
}
