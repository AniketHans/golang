package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func take(done <-chan int, primes <-chan int, n int) <-chan int { 
	//STAGE 3
	// here, we will only return n number of prime numbers using a taken channel
	taken := make(chan int)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-primes: // Adding a prime value to taken chan
			// the above code is similar to
			/*
			case val := <-primes:
				taken <- val
			*/
			}
		}
	}()
	return taken
}

func isPrime(n int) bool {
	// Checking if a number is prime or not
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(done <-chan int, stream <-chan int) <-chan int {
	// STAGE 2
	// Filtering out prime numbers from the stream of data received from STAGE 1
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randInt := <-stream:
				if isPrime(randInt) {
					primes <- randInt
				}
			}
		}
	}()
	return primes
}

func repeatFunc(done <-chan int, fn func() int) <-chan int {
	// STAGE 1
	// Generating a stream of randome Integers
	stream := make(chan int)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}


func fanIn(done <-chan int, primeFinderChannels... <-chan int) chan int{
	// STAGE for aggregating multiple STAGE 2 instances channel data into a single channel
	var wg sync.WaitGroup // We need waitgroups so that we can wait till all the data, from multiple channels, is transfered in one 
	fanInChan := make(chan int) // The single channel that will take data from multiple channels
	
	transfer := func (ch <-chan int){
		defer wg.Done() // Once the data transfer is completed or done channel is closed, we will mark the work as done
		for i := range ch{
			select{
			case <- done:
				return 
			case fanInChan <- i: // transfering data from passed ch channel into fanInChan
			}
		}
	}

	for _, ch := range primeFinderChannels{
		// Iterating over all the channels returned by all the instances of STAGE 2
		wg.Add(1)
		go transfer(ch) // Transfer will happen in separate thread to unblock the main thread
	}

	go func(){
		wg.Wait() // Wait, for all waitgrps to be done, will also be done on a separate thread to unblock the main thread
		close(fanInChan)
	}()

	return fanInChan // Since waiting for waitgrps for data transfer is being done in a separate thread, we can return the fanInChan to pass it to further stage
}

func main() {
	done := make(chan int)
	defer close(done)
	// //-------------------------------------------------- naive approach -----------------------------------------------------------
	// // This approach is slow because of single STAGE 2 instance
	// randnumFetcher := func() int { return rand.Intn(500000000)}
	// randomNumberStream := repeatFunc(done,randnumFetcher)
	// filteredPrimes := findPrimes(done, randomNumberStream) // If we have only one instance of this, we will get the results slowly
	// takenPrimes := take(done, filteredPrimes, 10)
	// for val := range takenPrimes{
	// 	fmt.Println(val)
	// }

	// // ------------------------------------------------- fan out and fan in --------------------------------------------------------
	// // This approach is faster because of multiple STAGE 2 instances
	randnumFetcher := func() int { return rand.Intn(500000000) }
	randomNumberStream := repeatFunc(done, randnumFetcher) //Stage 1, getting stream of random numbers

	// ----- fan out Stage 2 ------
	CPUCount := runtime.NumCPU() // getting the number of CPUs available in the system
	primeFinderChannels := make([]<-chan int, CPUCount) // buffered channel, for aggregating the result of all the STAGE 2 instance channels
	for i:=0; i<CPUCount; i++{
		primeFinderChannels[i] = findPrimes(done, randomNumberStream) // STAGE 2, here multiple, ~ No. of CPUs, STAGE 2 instances will be called
	}

	// ----- fan in Stage 2 ------
	fannedInChan := fanIn(done, primeFinderChannels...) // Agregating all data of all the channels returned by instaces of STAGE2 into a single channel
	
	// Stage 3
	takenPrimes := take(done, fannedInChan, 10) // STAGE 3, getting only n number of prime numbers from fanInChan
	for val := range takenPrimes {
		fmt.Println(val) // output
	}
}