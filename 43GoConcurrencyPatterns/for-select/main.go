package main

import (
	"fmt"
	"time"
)

// func main() {
// 	charChan := make(chan string, 3)
// 	chars := []string{"a", "b", "c"}

// 	for _, s := range chars {
// 		select {
// 		case charChan <- s:
// 		}

// 	}

// 	close(charChan) // Closing the channel

// 	for result := range charChan {
// 		// We can loop over a closed channel and receive the residual data left on the buffered channel
// 		// this is an unblocking code because loop will iterate till it reaches the channel closure
// 		fmt.Println(result)
// 	}

// 	// Output
// 	/*
// 		a
// 		b
// 		c
// 	*/

// }

// // ----------------------------- Infinite Goroutine --------------------------

// func main() {
// 	go func(){
// 		// infinitely looping go routine
// 		for {
// 			select{
// 			default:
// 				fmt.Println("DOING WORK....")
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Second*2)

// 	/*
// 		The above go routine is inifinitely going to print DOING WORK.... till the main() function returns
// 	*/

// }

// ----------------------------- Done Channel --------------------------

func main() {
	done := make(chan bool)
	go func(ch <-chan bool){
		// func (ch <-chan bool), is the syntax for sending a read only channel to a function
		// the function cannot write anything to the channel but only read data from it
		for {
			select{
			case <-ch: // 2. If some msg, in this case channel close msg, is received from the channel, this case will run and break the infinite loop. In this case, go routine will also stop
				return // If anything is received from the ch channel, the function/go routine will return
			default:
				fmt.Println("DOING WORK....")
			}
		}
	}(done)

	time.Sleep(time.Second*5)
	close(done) // 1. After the channel is closed, some channel close msg is send to the channel

	time.Sleep(time.Hour*2)
	
}