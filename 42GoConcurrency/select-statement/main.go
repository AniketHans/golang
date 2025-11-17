package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan int64)

	go func() {
		ch1 <- "data"
	}()

	go func() {
		ch2 <- 10
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Message received from channel ch1", msg1)
	case msg2 := <-ch2:
		fmt.Println("Message received from channel ch2", msg2)
	}

}