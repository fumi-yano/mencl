package entity

import (
   "time"
)

type Model struct {
   ID        uint `gorm:"primary_key" json:"id"`
   CreatedAt time.Time
   UpdatedAt time.Time
   DeletedAt *time.Time `sql:"index" json:"-"`
}

// User is user models property
type User struct {
  //  `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`フィールドを`User`モデルに注入
  Model
  Name string `json:"name"`
  Email string `json:"email"`
}
