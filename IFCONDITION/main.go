package main

import "fmt"

func main()  {
	// age := 10

	// if age > 10 {
	// 	fmt.Println("Boleh bermain game")
	// } else {
	// 	fmt.Println("Belum boleh bermain")
	// }

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}