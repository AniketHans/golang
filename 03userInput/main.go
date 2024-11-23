package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma of syntax
		//In golang, we dont have try catch. Go expects the problems and solutions to be treated as variables.
	input, _ := reader.ReadString('\n') //Igoring the error
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("The type of input is %T",input)
}