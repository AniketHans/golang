package main

import (
	"fmt"
	"sort"
)

func main() {
	// If you are using the below syntax, you need to initialize it as well
	var fruits = []string{"Apple", "Banana", "Tomato"}
	// var fruits = []string{} // This is also fine
	fmt.Printf("The type of fruits is, %T\n",fruits)

	//appending values
	fruits = append(fruits, "Mango","Peach")
	fmt.Printf("The fruits are, %v\n",fruits)

	//slicing the slice
	fmt.Println(fruits[1:3])

	// using make
	highscores := make([]int,4) // This will allocate a slice of with 4 slots init.
	highscores[0]=1
	highscores[1]=2
	highscores[2]=3
	highscores[3]=4

	// highscores[4]=5 //This will throw an error
	highscores = append(highscores, 5,6,7,8) 
	fmt.Println(highscores)
	// In the above case, Go behave diffrently. By default, highscores is a slice with 4 default values. 
	// But once more values are introduced in it using append, Go reallocates the memory to accomodate the new values.

	sort.Ints(highscores)
	fmt.Println(highscores)

 	//Removing a value from slice
	var courses = []string{"go","react","c++","js","python"}
	fmt.Println(courses)

	//remove the 2 index
	courses = append(courses[:2],courses[3:]... )
	fmt.Println(courses)


}