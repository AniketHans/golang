package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fmt.Println("Array is ",fruitList)
	fmt.Println("Length of the array is ",len(fruitList))

	fruitList[3] = "Tomato"
	fmt.Println("Array is ",fruitList)
	fmt.Println("Length of the array is ",len(fruitList))

	var vegList = [3]string{"Potato", "Brinjal", "Cauliflower"}
	fmt.Println("The veggies are ", vegList)
}