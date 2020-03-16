package user

import (
  _ "fmt"
  "github.com/gin-gonic/gin"
  "api/db"
  "api/entity"
  "gopkg.in/go-playground/validator.v9"
)

// Service procides user's behavior
// struct構造体について : https://dev.classmethod.jp/server-side/language/golang-5/
type Service struct{}

// User is alias of entity.User struct
// aliasについて : https://qiita.com/tenntenn/items/c3afc87a20d9f50998bb
type User entity.User

// GetAll is get all User
// レシーバーについて : https://teratail.com/questions/41829
func (s Service) GetAll() ([]User, error) {
    db := db.GetDB()
    var u []User

    if err := db.Find(&u).Error; err != nil {
        return nil, err
    }

    return u, nil
}

// CreateModel is create User model
func (s Service) CreateModel(c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User
    // https://gist.github.com/konojunya/453fa1e5f84058fd8cb8d7c0122348e1
    // Bindはparamsみたいなもの多分...

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    if err := db.Create(&u).Error; err != nil {
        return u, err
    }
    return u, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    return u, nil
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).First(&u).Error; err != nil {
        return u, err
    }

    if err := c.BindJSON(&u); err != nil {
        return u, err
    }

    db.Save(&u)

    return u, nil
}

// DeleteByID is delete a User
func (s Service) DeleteByID(id string) error {
    db := db.GetDB()
    var u User

    if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
        return err
    }

    return nil
}

// Validator宣言
var validate *validator.Validate

// BeforeSave User
func (u *User) BeforeSave() (err error) {
  validate = validator.New()
  err = validate.Struct(u)
  return err
}
