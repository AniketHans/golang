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


}