package main

import (
	"fmt"
)

func typ(ty interface{}) string {
	switch ty.(type) {
	case int:
		return "Interio"
	case string:
		return "Strings"
	case float32, float64:
		return "Real"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(typ("Wene"))
}
