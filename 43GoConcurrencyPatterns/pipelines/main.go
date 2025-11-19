package main

import "fmt"

func sliceToChannel(vals []int) chan int {
	out1 := make(chan int) // This is an unbuffered channel
	go func() { // This go routine will not block the sliceToChannel() and run on a separate thread
		for _, v := range vals {
			out1 <- v //Since out1 is unbuffered, so as soon as we put data in it, we need a consumer which is marked as Consumer1
			// The above line will wait until the value sent to out1 is consumed
		}
		close(out1) // After sending the values in the channel the channel is closed
	}()

	return out1 // The above go routine will start running on another thread and meanwhile the out1 chan is returned
}

func sq(ch <-chan int) chan int { // ch is the taken as a read only channel
	out2 := make(chan int) // This is an unbuffered channel
	go func() { // This go routine will not block the sq() and run on a separate thread
		for val := range ch { //Consumer1, also we will be able to iterate over the ch until it is closed. 
			// 	The looping will unblock the `out1 <- v` code as we are consuming the value from it by iterating and gets blocked itself until a new value is put to the channel for consumption
			out2 <- val * val // out2 here is also unbuffered so we need a consumer for the value we put in it marked as Consumer2
			// The above line will wait until the value sent to out2 is consumed
		}
		close(out2) // After sending the values in the channel the channel is closed
	}()
	return out2
}

func main() {
	values := []int{10, 20, 37, 47, 58}

	//Stage1 (Convert values slice to channel)
	out1 := sliceToChannel(values)

	//Stage2 (Square the values in the stage 1 returned channel)
	out2 := sq(out1)

	//Stage 3 (Print the square values returned from stage 2)
	for x := range out2 { //Consumer2,  also we will be able to iterate over the ch until it is closed
		// 	The looping will unblock the `out2 <- val * val` code as we are consuming the value from it by iterating and gets blocked itself until a new value is put to the channel for consumption
		fmt.Println(x)
	}

}


/*
	Here, above we have used unbuffered channels so the program is communicating synchronously.
	The go routines inside the sliceToChannel() and sq() will run on different threads while the functions returns the respective channels
	The returned channels will be consumed in further steps to unblock the go routines when they put some data on their respective channels
	The iteration over the channels will be performed till the channels are closed. If the channels are not closed, the iterators will keep waiting for value from the channel thus resulting in deadlock and error
*/
