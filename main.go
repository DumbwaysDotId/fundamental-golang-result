package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"dumbmerch/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

