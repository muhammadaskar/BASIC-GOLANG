package main

import "fmt"

func main()  {

	// for i:= 0; i < 100; i++{
	// 	fmt.Println("Saya Sedang Belajar Go Lang", i + 1)
	// }

	// i := 1
	// for i <= 100{
	// 	fmt.Println(i)
	// 	i++
	// }

	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("Index :", index, " letter :", string(letter))
	// }

	for i, letter := range title{
		if i % 2 == 0 {
			fmt.Print(string(letter))
		}

	}

	fmt.Println("")

	for _, letter := range title{
		switch string(letter){
		case "a" :
			fmt.Print(string(letter))
		case "i" :
			fmt.Print(string(letter))
		case "u" :
			fmt.Print(string(letter))
		case "e" :
			fmt.Print(string(letter))
		case "o" :
			fmt.Print(string(letter))
		}
	}
}