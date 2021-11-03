package main

import (
	"fmt"
	"time"
)

func isCousin(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func cousin(number int, ch chan int) {
	initial := 2

	for i := 0; i < number; i++ {
		for cousins := initial; ; cousins++ {
			if isCousin(cousins) {
				ch <- cousins
				initial = cousins + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(ch)
}

func main() {
	ch := make(chan int, 100)
	go cousin(cap(ch), ch)

	for cousins := range ch {

		fmt.Printf("%d ", cousins)

	}

	fmt.Println()
	for i := 0; i < 40; i++ {
		fmt.Print("-*")
	}
}
