package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channels in Golang")

	myCh := make(chan int) // declaring a channel of int type
	wg := &sync.WaitGroup{}
	buffCh := make(chan int, 1) // this is a buffered channel, here we can have max 2 values in channel
	// buffCh := make(chan int, 2) //here we can have max 3 (1+2) values in channel
	wg2 := &sync.WaitGroup{}

	// // pushing value into the channel
	// myCh <- 5 // note, we always have left pointing arrows '<-' in golang

	// // getting the value from the channel
	// fmt.Println(<-myCh)

	wg.Add(2) //Number of waitgroups being triggered
	go func(ch chan int, w *sync.WaitGroup){
		fmt.Println("Value from channel: ",<-ch)
		fmt.Println("Another value from channel: ",<-ch) //extra consumer as we are pushing 2 values below.
		defer wg.Done()
	}(myCh,wg)

	go func(ch chan int, w *sync.WaitGroup){
		fmt.Println("Adding to the channel")
		ch <- 5
		ch <- 6
		defer wg.Done()
	}(myCh,wg)

	//using buffered channel
	wg2.Add(2) //Number of waitgroups being triggered
	go func(ch chan int, w *sync.WaitGroup){
		fmt.Println("Value from Buffered channel: ",<-ch)
		defer wg2.Done()
	}(buffCh,wg2)

	go func(ch chan int, w *sync.WaitGroup){
		fmt.Println("Adding to the Buffered channel")
		ch <- 7
		ch <- 8
		//ch <- 9 //since buffCh can have max 2 values (7 & 8 from above), if we uncomment this, program will trow error
		defer wg2.Done()
	}(buffCh,wg2)



	wg.Wait()
	wg2.Wait()
}