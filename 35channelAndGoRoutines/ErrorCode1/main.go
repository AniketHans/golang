package main

import "fmt"

func main() {
	channel := make(chan string)
	go func1(channel)
	fmt.Println(<-channel) 
	fmt.Println(<-channel) // More recievers than senders. This is an error
}

func func1(c chan string){
	c <- "Hello"
}