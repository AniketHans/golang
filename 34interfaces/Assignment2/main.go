package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	// filedata, err := ioutil.ReadFile(filename)
	// if err!=nil{
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(filedata))

	// With using interfaces
	f, err := os.Open(filename)
	//since f *File type implements Reader interface
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout,f)

}