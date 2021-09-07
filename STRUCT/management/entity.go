package management

import "fmt"

type User struct {
	Id int
	FirstName string
	LastName string
	Email string
	Sex string
	IsActive bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s, Sex : %s", user.FirstName, user.LastName, user.Email, user.Sex)
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))

	fmt.Println("\nUsers Name")
	for _, user := range group.Users {
		fmt.Println(user.LastName)
	}
}

type PersegiPanjang struct {
	Panjang float32
	Lebar float32
}

func (pp PersegiPanjang) LuasPersegiPanjang() float32 {
	hasil := pp.Panjang * pp.Lebar
	return hasil
}