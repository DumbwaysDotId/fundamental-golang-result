# Delete data with ORM

### Repositories

> File: `repositories/users.go`

- Delete User data using `Delete` method

  ```go
  func (r *repository) DeleteUser(user models.User) (models.User, error) {
    err := r.db.Delete(&user).Error // Using Delete method

    return user, err
  }
  ```
