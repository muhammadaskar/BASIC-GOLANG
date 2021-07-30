package main

import "fmt"

func main()  {
	// var myMap map[string] int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["go"] = 9
	// myMap["java"] = 8
	// myMap["go"] = 10

	myMap := map[string] string{
		"ruby" : "is beautiful",
		"go" : "is super fast",
		"java" : "hype",
	}

	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)

	// for key, value := range myMap{
	// 	fmt.Println("Key : ", key, ", value : ", value)
	// }

	// fmt.Println("===============================")

	// delete(myMap, "ruby");

	// for key, value := range myMap{
	// 	fmt.Println("Key : ", key, ", value : ", value)
	// }
}