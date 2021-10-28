package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://backinsano.webnetwork.com.br/home")
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := Sum(7, 93)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%d \n", result)

	fmt.Println(res.Request.Proto)

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
}

func Sum(a int, b int) (int, error) {
	result := a + b
	if result > 100 {
		return 0, errors.New("Entradas maiores que 100")
	}

	return result, nil
}
