package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//We are hitting the JSServer backend apis
	fmt.Println("Get")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myUrl = "http://localhost:8000"
	response, err := http.Get(myUrl)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response Status",response.StatusCode)
	content,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	requestPayload := strings.NewReader(`
	{
		"name": "Naam",
		"address": "Earth"
	}
	`)

	response, err := http.Post(myurl,"application/json",requestPayload)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	decodedResponse, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(decodedResponse))

}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// making form data
	 data := url.Values{}
	 data.Add("Name","Aniket")
	 data.Add("Age","25")

	 response, err := http.PostForm(myurl, data)
	 if err!=nil{
		panic(err)
	 }

	 defer response.Body.Close()
	 content, _ := ioutil.ReadAll(response.Body)
	 fmt.Println(string(content))
}