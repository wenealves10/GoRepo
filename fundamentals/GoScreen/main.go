package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// Clears the screen
	screen.Clear()

	for {
		// Moves the cursor to the top left corner of the screen
		screen.MoveTopLeft()

		fmt.Println(time.Now().Unix())
		time.Sleep(time.Second)
	}

}
