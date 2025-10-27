package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    uint32
}

func main() {

	// structs
	aniket := User{"Aniket", "aniket@mail.com", true, 25}
	fmt.Println(aniket)

	// for getting more details
	fmt.Printf("The details are: %+v\n",aniket)
	fmt.Printf("Name is %v\n",aniket.Name)

	// using var
	var aniket2 User
	aniket2.Name="Aniket"
	aniket2.Email = "aniket@mail.com"
	fmt.Printf("The details are: %+v\n",aniket2)
	fmt.Printf("Name is %v\n",aniket2.Name)
}