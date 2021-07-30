package main

import (
	"fmt"
	"errors"
)

func main()  {
	// sentence := printMyResult("Saya sedang")
	// fmt.Println(sentence)

	// result := add(10, 20)
	// fmt.Println(result)

	// _, keliling := calculate(5, 2)
	// fmt.Println(luas)
	// fmt.Println(keliling)

	// Quiz
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	// result, errr := calculate(10, 2, "+")
	// result, errr := calculate(10, 2, "-")
	// result, errr := calculate(10, 2, "*")
	result, err := calculate(10, 2, "/")
	// result, errr := calculate(10, 2, "=")
	if (err != nil){
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar Golang"
	return newSentence
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}

// func calculate(panjang, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)

// 	return
// }

// QUIZ
func sum(param[] int) int {
	sum := 0

	for i := 0; i < len(param); i++{
		sum += param[i]
	}

	return sum
}

func calculate(a int, b int, c string) (result int, err error) {
	switch c{
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			err = errors.New("Unknown operation")
	}

	return
}
