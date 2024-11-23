package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	go greeter("world")
	time.Sleep(3*time.Second) // waited for the go routines to be completed.
}

func greeter(s string) {
	for i := 1; i < 6; i++ {
		fmt.Println(s)
	}
}