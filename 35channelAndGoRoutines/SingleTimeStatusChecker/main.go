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

	fmt.Println(<- c) // The main function will keep on waiting for the message as we have not closed the response, if we do resp.Body.Close() this line will also throw error as number of recievers is greater than senders.
	// Why the above extra reciever is not throwing any error : https://stackoverflow.com/questions/79460692/confusing-channel-behavior-in-go

	
	// for i:=0;i<len(ls);i++{
	// 	fmt.Println(<-c) // until we recieve the message from the channel, the execution of the for look blocks
	// }
}

func checkStatus(url string, ch chan string) {
	_, err := http.Get(url) // We are not consuming the respose here. Ideally we should atleast hold it to close the response like resp.Body.Close() otherwise it will lead to unexpected behaviour.
	if err!=nil{
		fmt.Println(url,"not working fine")
		ch <- "might be down!!" 
		return
	}
	fmt.Println(url, "is up!!!")
	ch <- "It is up"
}