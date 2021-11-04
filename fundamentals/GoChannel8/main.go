package main

import (
	"fmt"
	"time"
)

func toSpeak(person string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s talking %d", person, i)
		}
	}()

	return ch
}

func together(entry1, entry2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case t1 := <-entry1:
				ch <- t1
			case t2 := <-entry2:
				ch <- t2
			}
		}
	}()

	return ch
}

func main() {
	ch := together(
		toSpeak("Wene"),
		toSpeak("Ismael"),
	)

	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch, <-ch)
	}

}
