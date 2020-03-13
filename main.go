package main

import (
  "api/db"
  "api/server"
)

func main() {
  db.Init()
  server.Init()
  db.Close()
}

// func main() {
//     // DBMigrate(DBConnect())
//     r := gin.Default()
//     r.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "Hello world",
//         })
//     })
//
//     r.GET("/test", func(c *gin.Context) {
//         c.String(200, "Hello,World!")
//     })
//
//     r.GET("/test2", func(c *gin.Context) {
//         c.String(200, "Hello,World!")
//     })
//
//     r.GET("/test3", func(c *gin.Context) {
//         c.String(200, "Hello,World!")
//     })
//
//     db.Init()
//     r.Run(":8080")
//
//     db.Close()
// }

// テーブルの作成
// type User struct {
//     gorm.Model
//     NickName string `json:"nickName"`
// }
//
// func DBMigrate(db *gorm.DB) *gorm.DB {
//     db.DropTable(&User{})
//     // db.AutoMigrate(&User{})
//     return db
// }

// func DBConnect() *gorm.DB {
//     DBMS := "mysql"
//     USER := "root"
//     PASS := "password"
//     PORT := "3306"
//     HOST := "db"
//     DBNAME := "mencl_db"
//     CONNECT := USER + ":" + PASS + "@" + "tcp("+HOST+":"+PORT+")" + "/" + DBNAME + "?parseTime=true"
//     db, err := gorm.Open(DBMS, CONNECT)
//     if err != nil {
//         panic(err.Error())
//     }
//     return db
// }
