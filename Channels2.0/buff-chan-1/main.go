package main

import "fmt"

func func1() chan int {
	out := make(chan int,1)
	go func() {
		defer close(out)
		for x := range 5 {
			out <- x
			out <- x*10
			out <- x*100
		}
	}()
	return out
}

func main() {
	vals := func1()

	for val := range vals {  // This is a blocking call and will wait for the channel to send all the values
		fmt.Println("Received val:",val)
	}

	// The above code is similar to
	/*
		for {
			val, ok := <-vals // receive from channel
			if !ok {
				break // channel closed
			}
			fmt.Println(val)
		}
	*/
}