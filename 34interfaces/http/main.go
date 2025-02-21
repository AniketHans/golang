package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
func main() {
	resp, err := http.Get("https://www.google.com")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	// bs := make([]byte,99999)
	// bs := []byte{} // Not to send an empty byte slice.
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	// using custom struct that implements writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write (p []byte) (int,error){
	fmt.Println("*******************")
	fmt.Println(string(p))
	fmt.Println("*******************")
	return len(p),nil
}