package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6)+1
	fmt.Println("The number on the dice is, ",diceNumber)


	//Note: There is no need to use break explicitly in between cases here.
	// Means, like other languages fallthrough is not there in golang. Thus, no need to use break.
	switch diceNumber{
	case 1:
		fmt.Println("The dice value is 1 and you can open")
	case 6:
		fmt.Println("You can move the dice 6 spot and roll it again")
	default:
		fmt.Println("You can move the dice",diceNumber, "spot")
	}

	// Explicitly, involving fallthrough
	currentEngineeringYear := 3
	switch currentEngineeringYear{
	case 1:
		fmt.Println("You have to pass 1st year")
		fallthrough
	case 2:
		fmt.Println("You have to pass 2nd year")
		fallthrough
	case 3:
		fmt.Println("You have to pass 3rd year")
		fallthrough
	case 4:
		fmt.Println("You have to pass 4th year")
	default:
		fmt.Println("You have passed the Engineering")
	}
}