package main

import (
	"fmt"

	"github.com/wenealves10/htmlTitles"
)

func main() {
	t2 := htmlTitles.Titles("https://wenedev.site")
	fmt.Println(<-t2)
}
