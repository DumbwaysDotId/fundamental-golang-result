## Table of contents

- [Delete Query with Gorm](#delete-query-with-gorm)
  - [Repositories](#repositories)
  - [Handlers](#handlers)
  - [Routes](#routes)

# Delete Query with Gorm

### Repositories

> File: `repositories/users.go`

- Declare `DeleteUser` interface
  ```go
  type UserRepository interface {
    FindUsers() ([]models.User, error)
    GetUser(ID int) (models.User, error)
    CreateUser(user models.User) (models.User, error)
    UpdateUser(user models.User, ID int) (models.User, error)
    DeleteUser(user models.User, ID int) (models.User, error) // Write this code
  }
  ```
- Write `DeleteUser` function

  ```go
   // Write this code
  func (r *repository) DeleteUser(user models.User,ID int) (models.User, error) {
    err := r.db.Raw("DELETE FROM users WHERE id=?",ID).Scan(&user).Error

    return user, err
  }
  ```

### Handlers

> File: `handlers/users.go`

- Write `DeleteUser` function

  ```go
  // Write this code
  func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    user, err := h.UserRepository.GetUser(id)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    data, err := h.UserRepository.DeleteUser(user,id)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
    json.NewEncoder(w).Encode(response)
  }
  ```

### Routes

> File: `routes/users.go`

- Write `Delete User` route with `DELETE` method

  ```go
  func UserRoutes(r *mux.Router) {
    userRepository := repositories.RepositoryUser(mysql.DB)
    h := handlers.HandlerUser(userRepository)

    r.HandleFunc("/users", h.FindUsers).Methods("GET")
    r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
    r.HandleFunc("/user", h.CreateUser).Methods("POST")
    r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
    r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE") // Write this code
  }
  ```
