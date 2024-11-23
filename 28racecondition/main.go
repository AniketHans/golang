package main

import (
	"fmt"
	"sync"
)

func main() {

	var ints = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		defer wg.Done()
		fmt.Println("One")
		mut.Lock()
		ints = append(ints, 1)
		mut.Unlock()
	}(wg,mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		defer wg.Done()
		fmt.Println("Two")
		mut.Lock()
		ints = append(ints, 2)
		mut.Unlock()
	}(wg,mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		defer wg.Done()
		fmt.Println("Three")
		mut.Lock()
		ints = append(ints, 3)
		mut.Unlock()
	}(wg,mut)

	wg.Wait()

	fmt.Println(ints)

}