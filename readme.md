# Insert data with ORM

### Repositories

> File: `repositories/users.go`

- Update User data using `Save` method

  ```go
  func (r *repository) UpdateUser(user models.User) (models.User, error) {
    err := r.db.Save(&user).Error // Using Save method

    return user, err
  }
  ```
