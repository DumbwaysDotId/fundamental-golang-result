## Table of contents

- [Update Query with Gorm](#update-query-with-gorm)
  - [Repositories](#repositories)
  - [Handlers](#handlers)
  - [Routes](#routes)

# Update Query with Gorm

### Repositories

> File: `repositories/users.go`

- Declare `UpdateUser` interface
  ```go
  type UserRepository interface {
    FindUsers() ([]models.User, error)
    GetUser(ID int) (models.User, error)
    CreateUser(user models.User) (models.User, error)
    UpdateUser(user models.User, ID int) (models.User, error) // Write this code
  }
  ```
- Write `UpdateUser` function

  ```go
   // Write this code
  func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
    err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password,ID).Scan(&user).Error

    return user, err
  }
  ```

### Handlers

> File: `handlers/users.go`

- Write `UpdateUser` function

  ```go
  // Write this code
  func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    request := new(usersdto.UpdateUserRequest) //take pattern data submission
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
      json.NewEncoder(w).Encode(response)
      return
    }

    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    user := models.User{}

    if request.Name != "" {
      user.Name = request.Name
    }

    if request.Email != "" {
      user.Email = request.Email
    }

    if request.Password != "" {
      user.Password = request.Password
    }

    data, err := h.UserRepository.UpdateUser(user,id)
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

- Write `Update User` route with `PATCH` method

  ```go
  func UserRoutes(r *mux.Router) {
    userRepository := repositories.RepositoryUser(mysql.DB)
    h := handlers.HandlerUser(userRepository)

    r.HandleFunc("/users", h.FindUsers).Methods("GET")
    r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
    r.HandleFunc("/user", h.CreateUser).Methods("POST")
    r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH") // Write this code
  }
  ```
