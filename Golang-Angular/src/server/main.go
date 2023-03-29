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

	// User Account Signup
	r.GET("/signup", handlers.GetSignupHandler)
	r.POST("/signup", handlers.AddSignupHandler)
	
	// User Account Login
	r.GET("/login", handlers.GetLoginHandler)
	r.POST("/login", handlers.AddLoginHandler)


	// Weather API
	r.GET("/weather/forecast", handlers.GetWeatherForecastHandler)
	r.GET("/weather", handlers.GetWeatherHandler)

	// TODO: Limit to Admin?
	// Users REST API
	r.GET("/users", handlers.GetUsersHandler)
	r.POST("/users", handlers.AddUsersHandler)
	r.DELETE("/users/:uid", handlers.DeleteUsersHandler)

	// Schedules REST API
	r.GET("/schedule", handlers.GetSchedulesHandler)
	r.POST("/schedule", handlers.AddSchedulesHandler)
	r.DELETE("/schedule/:id", handlers.DeleteSchedulesHandler)

	// Course REST API
	r.GET("/course", handlers.GetClassesHandler)
	r.POST("/course", handlers.AddClassHandler)
	r.DELETE("/course/:id", handlers.DeleteClassHandler)
	// r.PUT("/course", handlers.EditClassHandler)

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
