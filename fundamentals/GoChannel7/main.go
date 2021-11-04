package main

import (
	"fmt"
	"time"

	"github.com/wenealves10/webScrapper"
)

func theFastest(url1, url2, url3 string) string {
	ch1 := webScrapper.Titles(url1)
	ch2 := webScrapper.Titles(url2)
	ch3 := webScrapper.Titles(url3)

	select {
	case t1 := <-ch1:
		return t1
	case t2 := <-ch2:
		return t2
	case t3 := <-ch3:
		return t3
	case <-time.After(2000 * time.Millisecond):
		return "all lose"
		// default:
		// 	return "no reply"
	}
}

func main() {
	champion := theFastest(
		"https://wenedev.site",
		"https://google.com",
		"https://amazon.com",
	)

	fmt.Println(champion)
}
