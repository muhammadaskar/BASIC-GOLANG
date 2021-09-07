package main

import "fmt"

type Student struct{
	ID int
	Name string
	GPA float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " Eng"
}

func main(){
	student := Student{1, "Muhammad Askar", 3.50}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}
