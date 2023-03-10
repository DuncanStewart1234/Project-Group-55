// This package is the main file of the application and what is called on the backend to run it
package main

import (
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/handlers"
	"github.com/gin-gonic/gin"
)

// main is used to set up the application and implements the API for the todo and notes lists
func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Notes REST API
	r.GET("/notes", handlers.GetNotesHandler)
	r.POST("/notes", handlers.AddNotesHandler)
	r.DELETE("/notes/:id", handlers.DeleteNotesHandler)
	r.PUT("/notes", handlers.EditNotesHandler)

	// Todo REST API
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
