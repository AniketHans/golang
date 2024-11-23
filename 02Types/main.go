package main

import "fmt"

//global variable
//JWTToken := "randomValue" //Invalid
var JWTToken string ="randomValue"  //Valid

//constants
const LoginToken string ="Logintoken" //They cant be changed. Also this is a public variable as first letter is uppercase

func main() {
	//String
	var username string = "Aniket"
	fmt.Println(username)
	fmt.Printf("The type of the variable is: %T \n",username)

	//Bool
	var isLoggedIn bool =true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the variable is: %T \n",isLoggedIn)

	//integers
	var smallVal uint8 =255
	fmt.Println(smallVal)
	fmt.Printf("The type of the variable is: %T \n",smallVal)

	// var smallVal2 uint8 =256 // This throws error
	// fmt.Println(smallVal2)
	// fmt.Printf("The type of the variable is: %T \n",smallVal2)


	//float
	var smallFloat float32 =255.35456467657567545
	fmt.Println(smallFloat)
	fmt.Printf("The type of the variable is: %T \n",smallFloat)

	//Implicit type
	var website = "www.google.com" //Automatically know, whats the type of variable.
	fmt.Println(website)
	fmt.Printf("The type of the variable is: %T \n",website)

	// No var style
	numberOfUsers:= 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("The type of the variable is: %T \n",numberOfUsers)

	// numberOfUsers = 3000.6767 //Error
	// fmt.Println(numberOfUsers)
	// fmt.Printf("The type of the variable is: %T \n",numberOfUsers)

	// Public variable
	fmt.Println(JWTToken)
	fmt.Printf("The type of the variable is: %T \n",JWTToken)

	// Const variable
	fmt.Println(LoginToken)
	fmt.Printf("The type of the variable is: %T \n",LoginToken)

}