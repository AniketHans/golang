package main

import "fmt"

func printMap(m map[string]uint){
	for key, value := range m{
		fmt.Printf("\nKey is %v and its value is %v",key,value)
	}
}

func main() {
	characterMap := map[string]uint{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(characterMap) //map[a:1 b:2 c:3]

	printMap(characterMap)
	/*
		Key is c and its value is 3
		Key is a and its value is 1
		Key is b and its value is 2
	*/
}