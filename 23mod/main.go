package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello duniya walon!!")

	r := mux.NewRouter()

	r.HandleFunc("/",greeter).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func greeter(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello Duniya!!! Kaise ho aap log?</h1>"))
}