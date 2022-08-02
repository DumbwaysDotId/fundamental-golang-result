package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Todos struct {
	Id string `json:"id"`
	Title string `json:"title"`
	IsDone bool `isDone:"isDone"`
}

var todos = []Todos{
	{
		Id: "1",
		Title: "Cuci tangan",
		IsDone: true,
	},
	{
		Id: "2",
		Title: "Jaga jarak",
		IsDone: false,
	},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", FindTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	r.HandleFunc("/todo", CreateTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}

func FindTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var todoData Todos
	var isGetTodo = false

	for _, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todoData = todo
		}
	}
	
	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoData)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	var data Todos

	json.NewDecoder(r.Body).Decode(&data)

	todos = append(todos, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	var data Todos
	var isGetTodo = false

	json.NewDecoder(r.Body).Decode(&data)

	for idx, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todos[idx] = data
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	var isGetTodo = false
	var index = 0

	for idx, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			index = idx
		}
	}

	if isGetTodo == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID: " + id + " not found")
		return
	}

	todos = append(todos[:index], todos[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ID: " + id + " delete success")
}