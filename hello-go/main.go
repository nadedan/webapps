package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/hello-go", func(c *gin.Context) {
        c.String(200, "Hello, World from Go!")
    })
    r.Run(":8080")
}