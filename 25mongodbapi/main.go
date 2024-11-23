package main

import (
	"fmt"
	"mongodbapi/router"
	"net/http"
)

func main() {
	r:= router.Router()
	fmt.Println("Server getting started....")
	http.ListenAndServe(":4002",r)
	fmt.Println("Listening at PORT 4002")
}