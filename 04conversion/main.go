package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating for pizza:")
	reader := bufio.NewReader(os.Stdin)

	rating, err := reader.ReadString('\n')
	if (err!=nil){
		panic("Error")
	}
	fmt.Printf("Thanks for rating, %s",rating)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if (err!=nil){
		fmt.Println(err)
	}
	fmt.Printf("The type of numRating is, %T",numRating)


}