package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is text is for file"
	filePath := "./demo.txt"
	file, err := os.Create("./demo.txt") //File creation is an OS opertaion
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	noOfBytesWritten, err := io.WriteString(file, content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The number of bytes written,",noOfBytesWritten)
	readFile(filePath)
}

func readFile(filePath string){
	data, err := ioutil.ReadFile(filePath) // Data read from a file or anywhere (internet), it is always read in bytes format
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(data))
}