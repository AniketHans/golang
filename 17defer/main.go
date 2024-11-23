package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func myDefer(){
	// defered statements are executed at the last, before the function is returned and in LIFO manner
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
	// here the stack will be:- [0, 1, 2, 3, 4] <-- TOP
}

func main() {
	// defered statements are executed at the last, before the function is returned and in LIFO manner
	defer myDefer()
	defer printMessage("Hello") 
	defer printMessage("Hi!!")
	defer printMessage("How are you??")
	fmt.Println("This will execute first!!") 			//Not defered
	defer fmt.Println("I am fine!!!") 					//Not defered
	fmt.Println("This will execute at the 2nd") 		//Not defered
	myDefer() 											//Not defered
	fmt.Println("This guy won..")
}

//here the stack will be like [myDefer(), Hello, Hi, How are you??, I am fine!!!] <-- TOP