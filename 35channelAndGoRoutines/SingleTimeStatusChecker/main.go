package main

import (
	"fmt"
	"net/http"
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
		// remember we only use go keyword in front of function calls
		go checkStatus(url,c)
	} 

	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)

	fmt.Println(<- c) // The main function will keep on waiting for the message 

	// for i:=0;i<len(ls);i++{
	// 	fmt.Println(<-c) // until we recieve the message from the channel, the execution of the for look blocks
	// }
}

func checkStatus(url string, ch chan string) {
	_, err := http.Get(url)
	if err!=nil{
		fmt.Println(url,"not working fine")
		ch <- "might be down!!" 
		return
	}
	fmt.Println(url, "is up!!!")
	ch <- "It is up"
}