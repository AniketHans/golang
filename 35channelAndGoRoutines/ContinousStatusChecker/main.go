package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	ls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}
	
	c := make(chan string)

	for _, url := range ls{
		go checkStatus(url,c)
	} 

	// for {
	// 	go checkStatus(<-c,c) // Here,we are recieving the status code.
	// }

	//New syntax
	// for val := range c{
	// 	go checkStatus(val,c)
	// }


	//Adding a pause 
	for val := range c{
		go func(v string, c chan string){
			time.Sleep(3 * time.Second)
			checkStatus(v,c)

		}(val,c)
	}
}

func checkStatus(url string, ch chan string) {
	_, err := http.Get(url)
	if err!=nil{
		fmt.Println(url,"not working fine")
		ch <- url
		return
	}
	fmt.Println(url, "is up!!!")
	ch <- url
}