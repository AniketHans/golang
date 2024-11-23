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

	// calling the method. The created object will call methods
	aniket.GetStatus()

	// The created object is passed as call by values.
	aniket.NewMail()
	fmt.Println("The object's email is:",aniket.Email) //The email remains the same as copy of aniket object is passed 

	// The below function accepts reference of the object
	aniket.ChangeAge()
	fmt.Println("The object's age is:",aniket.Age)


}

func (u User) GetStatus(){  //This is a method
	fmt.Println("Is User active:",u.Status)
}

func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("Changed email is:",u.Email)
}

func (u *User) ChangeAge(){
	u.Age=26
	fmt.Println("Changed age is:",u.Age)
}