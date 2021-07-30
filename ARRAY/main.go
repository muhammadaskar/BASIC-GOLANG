package main

import "fmt"

func main()  {
	// var lang[5] string

	// lang[0] = "Go"
	// lang[1] = "Php"
	// lang[2] = "Ruby"
	// lang[3] = "Java"
	// lang[4] = "C#"

	lang := [5] string{
		"Go",
		"Python",
		"Java",
		"Php",
		"C#",
	}

	fmt.Println(lang)
	fmt.Println(len(lang))
}