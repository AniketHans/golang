package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://learncoding.com:3000/learn/?coursename=golang&paymentid=bfujhvb1863786"

func main() {
	fmt.Println(myurl)

	// parsing the url
	result, err := url.Parse(myurl)
	if err!=nil{
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Printf("The type of qparams is, %T\n",qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])


	// Constructing the url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "hans.com:8000",
		Path: "/learn/",
		RawQuery: "user=aniket",
	}
	fmt.Println(partsOfUrl.String())


}