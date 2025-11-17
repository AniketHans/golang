package main

import (
	"fmt"
	"time"
)

// // ----------------------------------------------- Scenario 1 ------------------------------------------------------
// func main() {
// 	buffCh := make(chan string, 3)
// 	go func() {
// 		buffCh <- "a"
// 		buffCh <- "b"
// 		buffCh <- "c"
// 		buffCh <- "d"
// 	}()

// 	/*
// 	In above go routine, we are putting more values than the defined one into the buffered channel and also we are not cosuming any
// 	of the values, so code buffCh <- "d" will block the go routine.
// 	We will not get any error because go does not throw any error if any go routine, except the main go routine, get blocked
// 	*/
// }

// func main() {
// 	buffCh := make(chan string, 3)
// 	go func() {
// 		buffCh <- "a"
// 		fmt.Println("a appened in buffChan")
// 		buffCh <- "b"
// 		fmt.Println("b appened in buffChan")
// 		buffCh <- "c"
// 		fmt.Println("c appened in buffChan")
// 		buffCh <- "d" // the go routine will get blocked as the buffChan is full
// 		fmt.Println("d appened in buffChan")
// 	}()

// 	time.Sleep(time.Second*2) // Waiting so that the go routine gets executed

// 	/*
// 	In above go routine, we are putting more values than the defined one into the buffered channel and also we are not cosuming any
// 	of the values, so code buffCh <- "d" will block the go routine.
// 	We will not get any error because go does not throw any error if any go routine, except the main go routine, get blocked
// 	*/

// 	//Output

// 	/*
// 		a appened in buffChan
// 		b appened in buffChan
// 		c appened in buffChan
// 		(main go routine exited)
// 	*/

// }

// // ----------------------------------------------- Scenario 2 ------------------------------------------------------

// func main() {
// 	buffCh := make(chan string, 3)
// 	buffCh <- "a"
// 	fmt.Println("a appened in buffChan")
// 	buffCh <- "b"
// 	fmt.Println("b appened in buffChan")
// 	buffCh <- "c"
// 	fmt.Println("c appened in buffChan")
// 	buffCh <- "d" // the go routine will get blocked as the buffChan is full
// 	fmt.Println("d appened in buffChan")

// 	/*
// 	In above go routine, we are putting more values than the defined one into the buffered channel and also we are not cosuming any
// 	of the values, so code buffCh <- "d" will block the go routine.
// 	This time we will get error because the we are putting data in buffChan in main thread. In this case, main thread will get blocked and
// 	the go will throw error
// 	*/

// 	//Output with error

// 	/*
// 		a appened in buffChan
// 		b appened in buffChan
// 		c appened in buffChan
// 		fatal error: all goroutines are asleep - deadlock!
// 	*/

// }

// ----------------------------------------------- Scenario 3 ------------------------------------------------------

func main() {
	buffCh := make(chan string, 3)
	go func(){
		buffCh <- "a"
		fmt.Println("a appened in buffChan")
		buffCh <- "b"
		fmt.Println("b appened in buffChan")
		buffCh <- "c"
		fmt.Println("c appened in buffChan")
		buffCh <- "d" // the go routine will get blocked as the buffChan is full
		fmt.Println("d appened in buffChan")
	}()
	
	fmt.Println("One value from buffCh",<- buffCh) // Consuming one value from channel in FIFO order

	time.Sleep(time.Second*2)

	/*
	In above go routine, we are putting one more value than the defined into the buffered channel and but we are consuming one data from 
	the channel in another go routine. Thus the buffered channel will have the extra space in it so this time it will accomodate the extra 
	value in it.
	Thus there will be no error this time
	
		MAIN thread                                               Another thread
		-----------                                               ------------
		buffCh := make(chan string, 3)
		go routine triggered  ----------------------------------> go routine performing tasks in separate thread																							 |
		<- buffCh <-------------------------- value from buffCh from thread -----------------|
		Sleep(2 seconds)
	*/

	//Output (in no particular order)

	/*
		a appened in buffChan
		b appened in buffChan
		c appened in buffChan
		d appened in buffChan
		One value from buffCh a
	*/

}