package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	for i:=0;i<2;i++{
		go func1(i, channel)
	}
	fmt.Println(<- channel) 
	fmt.Println(<- channel)
	fmt.Println(<- channel)// No sender but a reciever. This is an error.
}

func func1(i int,c chan string){
	time.Sleep(2 * time.Second)
	c <- fmt.Sprintf("Hello %v",i)
}