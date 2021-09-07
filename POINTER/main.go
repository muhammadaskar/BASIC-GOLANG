package main

import "fmt"

func main(){
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
	
	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)

	number := 5
	fmt.Println("Nilai awal : ", number)

	change(&number, 100)

	fmt.Println("Nilai akhir", number)

}

func change(old *int, new int) int {
	*old = new
	return *old
}