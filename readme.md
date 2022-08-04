> This section we will using Database to store and manage user data (Create, Read, Update, Delete)

---

## Table of contents

- [Prepare](#prepare)
  - [Installation](#installation)
  - [Database](#database)
  - [Models](#models)
  - [Auto Migrate](#auto-migrate)
  - [Connection](#Connection)
  - [Data Transfer Object (DTO)](#data-transfer-object-dto)
- [Fetching Query with Gorm](#fetching-query-with-gorm)
  - [Repositories](#repositories)
  - [Handlers](#handlers)
  - [Routes](#routes)
  - [Root file main.go](#root-file-maingo)

# Prepare

### Installation:

- Gorm

  ```bash
  go get -u gorm.io/gorm
  ```

- MySql

  ```bash
  go get -u gorm.io/driver/mysql
  ```

- Validator

  ```go
  go get github.com/go-playground/validator/v10
  ```

### Database

- Create database named `dumbmerch`

### Models

- Create `models` folder, inside it Create `user.go` file, and write below code

  > File: `models/user.go`

  ```go
  package models

  import "time"

  // User model struct
  type User struct {
    ID          int			`json:"id"`
    Name 		    string		`json:"name" gorm:"type: varchar(255)"`
    Email		    string 		`json:"email" gorm:"type: varchar(255)"`
    Password 	  string		`json:"password" gorm:"type: varchar(255)"`
    CreatedAt 	time.Time	`json:"created_at"`
    UpdatedAt 	time.Time	`json:"updated_at"`
  }
  ```

### Auto Migrate

- Create `database` folder, inside it Create `migration.go` file, and write below code

  > File: `database/migration.go`

  ```go
  package database

  import (
    "dumbmerch/models"
    "dumbmerch/pkg/mysql"
    "fmt"
  )

  // Automatic Migration if Running App
  func RunMigration() {
    err := mysql.DB.AutoMigrate(&models.User{})

    if err != nil {
      fmt.Println(err)
      panic("Migration Failed")
    }

    fmt.Println("Migration Success")
  }
  ```

### Connection

- Create `pkg` folder, inside it Create `mysql` folder, inside it Create `mysql.go` file, and write below code

  > File: `pkg/mysql/mysql.go`

  ```go
  package mysql

  import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
  )

  var DB *gorm.DB

  // Connection Database
  func DatabaseInit() {
    var err error
    dsn := "{USER}:{PASSWORD}@tcp({HOST}:{POST})/{DATABASE}?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
      panic(err)
    }

    fmt.Println("Connected to Database")
  }
  ```

### Data Transfer Object (DTO)

- Create `dto` folder, inside it create `result` & `users` folder.

  > Folder: `dto/result`

  > Folder: `dto/users`

- Inside `dto/result` folder, create `result.go` file, and write this below code

  > File: `dto/result/result.go`

  ```go
  package dto

  type SuccessResult struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
  }

  type ErrorResult struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
  }
  ```

- Inside `dto/users` folder, create `user_request.go` file, and write this below code

  > File: `dto/users/user_request.go`

  ```go
  package usersdto

  type CreateUserRequest struct {
    Name     string `json:"name" form:"name" validate:"required"`
    Email    string `json:"email" form:"email" validate:"required"`
    Password string `json:"password" form:"password" validate:"required"`
  }

  type UpdateUserRequest struct {
    Name     string `json:"name" form:"name"`
    Email    string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
  }
  ```

- Inside `dto/users` folder, create `user_response.go` file, and write this below code

  > File: `dto/users/user_response.go`

  ```go
  package usersdto

  type UserResponse struct {
    ID       int    `json:"id"`
    Name     string `json:"name" form:"name" validate:"required"`
    Email    string `json:"email" form:"email" validate:"required"`
    Password string `json:"password" form:"password" validate:"required"`
  }
  ```

# Fetching Query with Gorm

### Repositories

- Create `repositories` folder, inside it create `users.go` file, and write this below code

  > File: `repositories/users.go`

  ```go
  package repositories

  import (
    "dumbmerch/models"
    "gorm.io/gorm"
  )

  type UserRepository interface {
    FindUsers() ([]models.User, error)
    GetUser(ID int) (models.User, error)
  }

  type repository struct {
    db *gorm.DB
  }

  func RepositoryUser(db *gorm.DB) *repository {
    return &repository{db}
  }

  func (r *repository) FindUsers() ([]models.User, error) {
    var users []models.User
    err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

    return users, err
  }

  func (r *repository) GetUser(ID int) (models.User, error) {
    var user models.User
    err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

    return user, err
  }
  ```

### Handlers

- On `handlers` folder, create `users.go` file, and write this below code

  > File: `handlers/users.go`

  ```go
  package handlers

  import (
    dto "dumbmerch/dto/result"
    usersdto "dumbmerch/dto/users"
    "dumbmerch/models"
    "dumbmerch/repositories"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
  )

  type handler struct {
    UserRepository repositories.UserRepository
  }

  func HandlerUser(UserRepository repositories.UserRepository) *handler {
    return &handler{UserRepository}
  }

  func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    users, err := h.UserRepository.FindUsers()
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
    }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Code: http.StatusOK, Data: users}
    json.NewEncoder(w).Encode(response)
  }

  func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    user, err := h.UserRepository.GetUser(id)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
    json.NewEncoder(w).Encode(response)
  }

  func convertResponse(u models.User) usersdto.UserResponse {
    return usersdto.UserResponse{
      ID:       u.ID,
      Name:     u.Name,
      Email:    u.Email,
      Password: u.Password,
    }
  }
  ```

### Routes

- On `routes` folder, create `users.go`, and write this below code

  > File: `routes/users.go`

  ```go
  package routes

  import (
    "dumbmerch/handlers"
    "dumbmerch/pkg/mysql"
    "dumbmerch/repositories"
    "github.com/gorilla/mux"
  )

  func UserRoutes(r *mux.Router) {
    userRepository := repositories.RepositoryUser(mysql.DB)
    h := handlers.HandlerUser(userRepository)

    r.HandleFunc("/users", h.FindUsers).Methods("GET")
    r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
  }
  ```

- Modify `routes.go` file, like this below code

  > File: `routes/routes.go`

  ```go
  package routes

  import (
    "github.com/gorilla/mux"
  )

  func RouteInit(r *mux.Router) {
    TodoRoutes(r)
    UserRoutes(r)
  }
  ```

### Root file `main.go`

Modify `main.go` file, adding `Initial Database` and Running `Auto Migration`

```go
package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
```
