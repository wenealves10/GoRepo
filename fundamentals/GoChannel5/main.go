package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titles(urls ...string) <-chan string {
	channel := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\/title>`)
			channel <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return channel
}

func main() {
	t1 := titles("https://wenedev.site", "https://www.google.com.br")
	fmt.Println(<-t1, "\n", <-t1)
}
