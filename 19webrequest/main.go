package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	response, err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("The response type, %T\n",response)
	fmt.Println("Status:", response.Status)
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(dataBytes))

	defer response.Body.Close() // This is the callers responsibility to close the connection
}