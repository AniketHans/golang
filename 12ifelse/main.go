package main

import "fmt"

func main() {
	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount >= 10 && loginCount <= 20 {
		result = "Something's fishy"
	} else {
		result = "Fishy user"
	}

	fmt.Println(result)

	//something new

	if x:=10; x<20{
		fmt.Println("The number is less than 20")
	}else{
		fmt.Println("The number if equal to greater than 20")
	}
}