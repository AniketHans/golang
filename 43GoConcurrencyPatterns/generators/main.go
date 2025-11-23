package main

import (
	"fmt"
	"math/rand"
)

//Note: Since we are using unbuffered channel here, the data exchange between go routines is happening synchronously

func repeatFunc[T any](done <-chan int, fn func() T) <-chan T {
	// This repeatFunc function accepts a function as argument and calls that function repeatedly. 
	// The passed function will return a value of type T (say int, string, struct etc) and the value will be send to the stream chan 
	stream := make(chan T)
	go func() {
		// The go routine will run separately on a different thread
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn(): // Repeatedly calling the fn() recieved in the parameter
			}
		}
	}()
	return stream
}

func main() {
	done := make(chan int)
	defer close(done)
	randnumFetcher := func() int { return rand.Intn(500000000)}
	for val := range repeatFunc(done,randnumFetcher){
		// Here, we are looping over the chan returned by the repeatFunc() which contains the values returned by the randnumFetcher()
		fmt.Println(val)
	}
}

/*
	Here, we are using a go routines to continously generate stream of data.
	In real life, we might be getting continous stream of data from some IOT sensors or some website 
	which we may have to process further to make the more sense out of the data
*/