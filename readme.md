## Table of contents

- [Insert Query with Gorm](#insert-query-with-gorm)
  - [Repositories](#repositories)
  - [Handlers](#handlers)
  - [Routes](#routes)

# Insert Query with Gorm

### Repositories

> File: `repositories/users.go`

- Import `time`

  ```go
  import (
    "dumbmerch/models"
    "time"
    "gorm.io/gorm"
  )
  ```

- Declare `CreateUser` interface
  ```go
  type UserRepository interface {
    FindUsers() ([]models.User, error)
    GetUser(ID int) (models.User, error)
    CreateUser(user models.User) (models.User, error) // Write this code
  }
  ```
- Write `CreateUser` function

  ```go
   // Write this code
  func (r *repository) CreateUser(user models.User) (models.User, error) {
    err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)",user.Name,user.Email, user.Password, time.Now(), time.Now()).Error

    return user, err
  }
  ```

### Handlers

> File: `handlers/users.go`

- Import `Validator`

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
    "github.com/go-playground/validator/v10"  // Write this code
    "github.com/gorilla/mux"
  )
  ```

- Write `CreateUser` function

  ```go
   // Write this code
  func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    request := new(usersdto.CreateUserRequest)
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    validation := validator.New()
    err := validation.Struct(request)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    // data form pattern submit to pattern entity db user
    user := models.User{
      Name:     request.Name,
      Email:    request.Email,
      Password: request.Password,
    }

    data, err := h.UserRepository.CreateUser(user)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(w).Encode(err.Error())
    }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
    json.NewEncoder(w).Encode(response)
  }
  ```

### Routes

> File: `routes/users.go`

- Write `Create User` route with `POST` method

  ```go
  func UserRoutes(r *mux.Router) {
    userRepository := repositories.RepositoryUser(mysql.DB)
    h := handlers.HandlerUser(userRepository)

    r.HandleFunc("/users", h.FindUsers).Methods("GET")
    r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
    r.HandleFunc("/user", h.CreateUser).Methods("POST")  // Write this code
  }
  ```
