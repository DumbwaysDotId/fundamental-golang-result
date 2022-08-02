# Group routes

Group Routes are needed in API development to differentiate a route for API or for standard website link.

- Create `routes` folder source inside it have `routes.go` and `todos.go` file

- Create `handlers` folder source inside it have `todo.go` file

---

- On `routes/routes.go` file, declare Grouping Function for all Route

  ```go
  package routes

  import (
  	"github.com/gorilla/mux"
  )

  func RouteInit(r *mux.Router) {
  	TodoRoutes(r)
  }

  ```

- On `routes/todos.go` file, declare route and handler

  > File: `routes/todos.go`

  ```go
  package routes

  import (
  	"dumbmerch/handlers"
  	"github.com/gorilla/mux"
  )

  func TodoRoutes(r *mux.Router) {

  	r.HandleFunc("/todos", handlers.FindTodos).Methods("GET")
  	r.HandleFunc("/todo/{id}", handlers.GetTodo).Methods("GET")
  	r.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
  	r.HandleFunc("/todo/{id}", handlers.UpdateTodo).Methods("PATCH")
  	r.HandleFunc("/todo/{id}", handlers.DeleteTodo).Methods("DELETE")
  }

  ```

- On `handlers/todo.go` file, declare `struct`, `dummy data`, and the handlers function

  ```go
  package handlers

  import (
  	"net/http"
  	"github.com/gorilla/mux"
  	"encoding/json"
  )

  // Todos struct
  type Todos struct {
  	Id string `json:"id"`
  	Title string `json:"title"`
  	IsDone bool `isDone:"isDone"`
  }

  // Todos dummy data
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

  ```go
  // Get All Todo
  func FindTodos(w http.ResponseWriter, r *http.Request){
  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(http.StatusOK)
  	json.NewEncoder(w).Encode(todos)
  }

  ```

  ```go
  // Get Todo by Id
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

  ```go
  // Create Todo
  func CreateTodo(w http.ResponseWriter, r *http.Request){
  	var data Todos

  	json.NewDecoder(r.Body).Decode(&data)

  	todos = append(todos, data)

  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(http.StatusOK)
  	json.NewEncoder(w).Encode(todos)
  }
  ```

  ```go
  // Update Todo
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

  ```go
  // Delete Todo
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
