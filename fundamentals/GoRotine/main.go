package main

import (
	"fmt"
	"time"
)

func whatsUp(name, text string, qt int) {
	for i := 0; i < qt; i++ {
		fmt.Printf("%d) %s %s\n", i+1, name, text)
		time.Sleep(time.Second)
	}
}

func main() {
	// whatsUp("Wene", "Eae Insano", 10)
	// whatsUp("Ismael", "Eae Loouco", 20)

	go whatsUp("Wene", "Eae Insano", 10)
	go whatsUp("Ismael", "Eae Loouco", 20)
	whatsUp("Lucas", "Eae Loouco", 20)

}
