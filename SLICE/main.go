package main

import "fmt"

func main()  {
	var gamingConsole[] string

	gamingConsole = append(gamingConsole, "PS 3")
	gamingConsole = append(gamingConsole, "Nintendo")
	gamingConsole = append(gamingConsole, "Xbox One")

	// gamingConsole := []string {
	// 	"PS 4",
	// 	"Nintendo",
	// 	"Xbox",
	// }

	for _, console := range gamingConsole{
		fmt.Println(console)
	}

}