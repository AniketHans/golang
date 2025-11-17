package main

import (
	"fmt"
	"time"
)

func someFunc(num int64) {
	fmt.Println(num)
}

func main() {
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)

	time.Sleep(time.Second*2) // this time.sleep will only prevent our main() function to return before the go routines finish there work
	// Note: the above time.sleep is not acting as way of joining the go routines back to main(). It is just holding the main() from exiting
	fmt.Println("Hi")
}