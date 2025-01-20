package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    r.POST("/post", func(c *gin.Context) {
        c.String(http.StatusOK, "POST request")
    })

    r.Run() // Listen and serve on 0.0.0.0:8080
}
