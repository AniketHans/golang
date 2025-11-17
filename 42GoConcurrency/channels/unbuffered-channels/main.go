package main

import "fmt"

func main() {
	ch := make(chan string) // ch is an unbuffered channel

	go func() {
		// close(ch)
		ch <- "Data from go routine" // This is the syntax for sending some data to the channel
	}()

	// Here below, the main() go routine will wait for the message from the channel
	msg := <-ch // syntax for reading data from the channel. This code will act as the join code of go routine to the main()
	/*
		Note: unbuffered channels are blocking in nature. If you are reading data from an unbuffered channel, whether it is the main thread
		or any other go routine, the go routine's execution will be holded at the line of code where it is reading from the channel. It will
		only be unblocked once it recieves some value from the channel
		Thus while using unbuffered channels and reading data from the channel, make sure you have a someone sending the data to it.
		You cannot send or recieve data in the same function from the same unbuffered channel. This is because as soon as you send data to the
		channel, you have to consume it. But if you are doing both sending and reading data from channel in the same function, you wont be able
		to achieve the above condition due to sync execution of code within the same function
	*/
	fmt.Printf("%v",msg)
}