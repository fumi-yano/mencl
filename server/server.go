package server

import (
    "github.com/gin-gonic/gin"
    "api/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run(":8080")
}

func router() *gin.Engine {
    r := gin.Default()

    v := r.Group("/v1")
    u := v.Group("/users")
    {
        ctrl := user.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    return r
}
