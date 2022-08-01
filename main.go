package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	// On Terminal/Command Propt
	fmt.Println("Hello World!")

	// On http (API)
	r := mux.NewRouter()

	r.HandleFunc("/", helloWorld ).Methods("GET")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}


func helloWorld(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}