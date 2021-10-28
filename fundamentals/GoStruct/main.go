package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// type ClientName string
// type ClientEmail string
// type ClientCPF int

type Client struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

type ClientInternacial struct {
	Client
	Contry string `json:"contry"`
}

func (c Client) Show() {
	fmt.Println("Name: ", c.Name)
}

func main() {
	clientOne := Client{
		Name:  "Wene Alves",
		Year:  19,
		Email: "wene.alves@com.br",
		CPF:   1111111111111,
	}

	clientTwo := Client{"Wene Alves 2", 19, "wene@com.br", 2222222222}

	fmt.Println(clientOne)
	fmt.Println(clientTwo)
	fmt.Printf("%s %d %s %d\n", clientOne.Name, clientOne.Year, clientOne.Email, clientOne.CPF)

	clientInternacial := ClientInternacial{
		Client: Client{
			Name:  "Wene Alves Inter",
			Year:  19,
			Email: "wene.alves@com.br",
			CPF:   1111111111111,
		},
		Contry: "EUA",
	}

	fmt.Printf("%s %d %s %d %s\n", clientInternacial.Name, clientInternacial.Year, clientInternacial.Email, clientInternacial.CPF, clientInternacial.Contry)

	clientOne.Show()
	clientTwo.Show()
	clientInternacial.Show()

	jsonClien, err := json.Marshal(clientInternacial)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(jsonClien))

	jsonClientInter := `{"name":"Wene Alves Inter","year":19,"email":"wene.alves@com.br","cpf":1111111111111,"contry":"EUA"}`

	clientInternacialTwo := ClientInternacial{}

	json.Unmarshal([]byte(jsonClientInter), &clientInternacialTwo)

	fmt.Println(clientInternacialTwo)

}
