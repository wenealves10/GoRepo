package main

import "fmt"

func main() {
	fmt.Print("Hello ")
	fmt.Print("World")
	fmt.Println(" Wene")
	fmt.Print("Alves")

	valor := 3.1242

	xs := fmt.Sprint(valor)
	fmt.Println("Valor: " + xs)
	fmt.Println("Valor: ", valor)

	fmt.Printf("Valor: %.2f", valor)

	a := 1
	b := 1.99999
	c := false
	d := "Hellor"

	fmt.Printf("Valores: %d %f %.2f %t %s", a, b, b, c, d)
	fmt.Printf("Valores: %v %v %v %v %v", a, b, b, c, d)

}
