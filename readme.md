# Insert data with ORM

### Repositories

> File: `repositories/users.go`

- Add User data using `Create` method

  ```go
  func (r *repository) CreateUser(user models.User) (models.User, error) {
    err := r.db.Create(&user).Error

    return user, err
  }
  ```
