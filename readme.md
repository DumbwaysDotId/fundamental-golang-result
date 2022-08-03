> This section we will using Database to store data

---

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

### Database, Models, Auto Migrate, and Connection

- Create database named `dumbmerch`

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

- Create `database` folder, inside it Create `migration.go` file, and write below code

  > File: `datbase/migration.go`

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
