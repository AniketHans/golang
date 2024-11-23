package main

import "fmt"

func main() {
	// maps
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["PY"])

	//deleting from maps
	delete(languages,"RB")
	fmt.Println(languages)

	//This will return the default value of type
	fmt.Println("The value of KT is, ",languages["KT"])

	//loops on maps
	for key, value := range languages{
		fmt.Println(key," ----> ",value)
	}

	// using var
	var numberToAlpha = map[int]string{}
	numberToAlpha[1]="A"
	fmt.Println(numberToAlpha)

}