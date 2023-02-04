package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/handlers"
)

func main() {
    r := gin.Default()
    r.Use(CORSMiddleware())

    r.GET("/todo", handlers.GetTodoListHandler)
    r.POST("/todo", handlers.AddTodoHandler)
    r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
    r.PUT("/todo", handlers.CompleteTodoHandler)

    err := r.Run(":3000")
    if err != nil {
        panic(err)
    }
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}