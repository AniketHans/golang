package main

import "fmt"

func main() {
	var ptr *uint32
	fmt.Printf("The value of pointer is : %v\n",ptr)
	fmt.Println("The address of the pointer itself is: ",&ptr)

	number:= 56
	var ptr1 *int = &number
	fmt.Printf("The value of pointer is : %v\n",ptr1)
	fmt.Println("The address of the pointer itself is: ",&ptr1)
	fmt.Printf("The value stored at the pointer is : %v\n",*ptr1)

	*ptr1 = *ptr1*2
	fmt.Println("The new value of number is, ",number)
}