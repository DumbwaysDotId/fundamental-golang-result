# Routing

Routing in `gorilla/mux` is pretty straightforward. Make sure you already understand the core concept of REST API like using GET/POST/PATCH/DELETE etc.

`API` or `Application Programming Interface` is an interface that can connect one application with another application. Thus, the API acts as an intermediary between different applications, either within the same platform or across platforms.

> File: `main.go`

- package name

  ```go
  package main
  ```

- import package

  ```go
  import (
  	"fmt"
  	"net/http"
  	"github.com/gorilla/mux"
  	"encoding/json"
  )
  ```

- Struct

  ```go
  type Todos struct {
  	Id string `json:"id"`
  	Title string `json:"title"`
  	IsDone bool `isDone:"isDone"`
  }
  ```

- Dummy Data

  ```go
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
  ```

- Main function for Declare Route

  ```go
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
  ```

- Get all Todo data

  ```go
  func FindTodos(w http.ResponseWriter, r *http.Request){
  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(http.StatusOK)
  	json.NewEncoder(w).Encode(todos)
  }
  ```

- Get Todo data by Id

  ```go
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
  ```

- Create Todo Data

  ```go
  func CreateTodo(w http.ResponseWriter, r *http.Request){
  	var data Todos

  	json.NewDecoder(r.Body).Decode(&data)

  	todos = append(todos, data)

  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(http.StatusOK)
  	json.NewEncoder(w).Encode(todos)
  }
  ```

- Update Todo Data

  ```go
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
  ```

- Delete Todo data

  ```go
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
  ```
