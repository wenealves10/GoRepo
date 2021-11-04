package main

import (
	"fmt"

	"github.com/wenealves10/webScrapper"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func together(prohibited, prohibited2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(prohibited, c)
	go forward(prohibited2, c)

	return c
}

func main() {
	ch := together(webScrapper.Titles("https://google.com"),
		webScrapper.Titles("https://wenedev.site"))
	fmt.Println(<-ch + " <----> " + <-ch)
}
