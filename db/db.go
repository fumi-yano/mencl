package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "api/entity"
)

var (
  db  *gorm.DB
  err error
)

// Init is initialize db from main function
func Init() {
  DBMS := "mysql"
  USER := "root"
  PASS := "password"
  PORT := "3306"
  HOST := "db"
  DBNAME := "mencl_db"
  CONNECT := USER + ":" + PASS + "@" + "tcp("+HOST+":"+PORT+")" + "/" + DBNAME + "?parseTime=true"
  db, err = gorm.Open(DBMS, CONNECT)
  if err != nil {
    panic(err)
  }

  autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}


func autoMigration() {
  // db.Delete(&entity.User{})
  db.AutoMigrate(&entity.User{})
}
