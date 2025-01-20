package main

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func main() {
    e := echo.New()
    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.POST("/post", func(c echo.Context) error {
        return c.String(http.StatusOK, "POST request")
    })

    e.Start(":8080") // Start the server on port 8080
}
