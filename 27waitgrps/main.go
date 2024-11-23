package main

import (
	"fmt"
	"net/http"
	"sync"
)

var strings = []string{"urls"}

var wg sync.WaitGroup //pointer type
var mut sync.Mutex // mutex pointer

func main() {
	websiteList :=[]string{
		"https://www.google.com","https://www.amazon.com","https://www.flipkart.in","https://www.facebook.com",
	}

	for _,web := range websiteList{
		go getStatusCode(web)
		wg.Add(1) // We add the number of go routines being called each time.
	}

	wg.Wait() // This wait menthod always go at the end of main method
	fmt.Println(strings)

}

func getStatusCode(endpoint string) {
	defer wg.Done() // This tells that this goroutine is done.
	res, err := http.Get(endpoint)
	if err!=nil{
		fmt.Println("Oops in endpoints!!!")
	}else{
		mut.Lock()
		strings = append(strings, endpoint)
		mut.Unlock()

		fmt.Printf("The status code for %s : %d\n",endpoint,res.StatusCode)
	}
	
}