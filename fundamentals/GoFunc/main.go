package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Model string
}

func (c Car) walk() {
	fmt.Println(c.Name, "Walking...")
}

func main() {

	car := Car{
		Name:  "Fiat uno",
		Year:  2002,
		Model: "Gol",
	}

	car.walk()

	var resOne = sumOne(1, 2)
	fmt.Println(resOne)

	resTwo, str := sumTwo(10, 2)
	fmt.Println(resTwo, str)

	var resThree = sumThree(10, 10)
	fmt.Println(resThree)

	resFor := sumFor(1, 2, 3, 4, 40)
	fmt.Println(resFor)

	resultFine := func(x ...int) func() int {
		result := 0

		for _, value := range x {
			result += value
		}

		return func() int {
			return result * result
		}
	}

	fmt.Println(resultFine(1, 2, 3)())
}

func sumOne(a int, b int) int {
	return a + b
}

func sumTwo(a int, b int) (int, string) {
	return a + b, "Done!"
}

func sumThree(a int, b int) (result int) {
	result = a + b
	return
}

func sumFor(x ...int) int {
	result := 0

	for _, value := range x {
		result += value
	}

	return result
}
