package main

import (
	"fmt"
	"STRUCT/management"
)





func main()  {
	user := management.User{1, "Muhammad", "Askar", "askar@gmail.com", "Male", true}

	result := user.Display()
	fmt.Println(result)
	
	
	user2 := management.User{2, "Muhammad", "Aswad", "aswad@gmail.com", "Male" , true}
	
	result2 := user2.Display()
	fmt.Println(result2)

	fmt.Println("---------------------")

	pp := management.PersegiPanjang{5, 3}
	// hasil := pp.luas
	fmt.Println("Hasil Luas Persegi Panjang : ", pp.LuasPersegiPanjang())

	// users := []management.User{user, user2}
	// group := management.Group{"Gamer", user, users, true}

	// group.DisplayGroup()

}