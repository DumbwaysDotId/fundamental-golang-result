## Table of contents

- [Fetching data with ORM](#fetching-data-with-orm)
  - [Repositories](#repositories)
  - [Handlers](#handlers)
  - [Routes](#routes)

# Fetching data with ORM

### Repositories

> File: `repositories/users.go`

- Find User all data using `Find` method

  ```go
  func (r *repository) FindUsers() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error // Using Find method

    return users, err
  }
  ```

- Get User data by ID using `First` method

  ```go
  func (r *repository) GetUser(ID int) (models.User, error) {
    var user models.User
    err := r.db.First(&user, ID).Error // Using First method

    return user, err
  }
  ```
