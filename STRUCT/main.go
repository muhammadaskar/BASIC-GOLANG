package main

import "fmt"

type User struct {
	id int
	firstName string
	lastName string
	email string
	isActive bool
}

func main()  {
	user := User{}
	user.id = 1
	user.firstName = "Muhammad"
	user.lastName = "Askar"
	user.email = "user@email.com"
	user.isActive = true

	user2 := User{
		id : 2,
		firstName : "Muhammad",
		lastName : "Aswad",
		email : "user2@gmail.com",
		isActive : true,
	}

	user3 := User{
		3,
		"Maulidya",
		"Lidya",
		"user3@gmail.com",
		true,
	}


	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)
}