// This package implements and maintains handler functions for our application
package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
	"github.com/gin-gonic/gin"
)

// To Do Handlers
// GetTodoListHandler is a handler to request the todo list from the Todo package
func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todo.Get())
}

// AddTodoHandler is used to add to the todo list from the Todo package
func AddTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}

	id, _ := todo.Add(todoItem.Message)
	c.JSON(statusCode, gin.H{"id": id})
}

// DeleteTodoHandler is used to delete an item from the todo list
func DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")
	if err := todo.Delete(todoID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// CompleteTodoHandler is used to mark a todo list item as complete
func CompleteTodoHandler(c *gin.Context) {
	todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if todo.Complete(todoItem.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

// Notes Handlers
// GetNotesHandler is used to request the notes list from the Notes package
func GetNotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, notes.Get())
}

// AddNotesHandler is used to add a note to the notes list from the Notes package
func AddNotesHandler(c *gin.Context) {
	notesItem, statusCode, err := convertHTTPBodyToNote(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": notes.Add(notesItem.Title, notesItem.Message)})
}

// DeleteNotesHandler is used to delete a note from the notes list
func DeleteNotesHandler(c *gin.Context) {
	noteID := c.Param("id")
	if err := notes.Delete(noteID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "peration Completed Successfully!")
}

// EditNotesHandler is used to edit a note from the notes list
func EditNotesHandler(c *gin.Context) {
	noteItem, statusCode, err := convertHTTPBodyToNote(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if notes.Edit(noteItem.ID, noteItem.Message) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Operation Completed Successfully!")
}

// Schedule + Maps Handlers

// User Login Handlers
// convertHTTPBodyToTodo is used to convert user input on the frontend for use with the backend Todo List
func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todo.Todo, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return todo.Todo{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToTodo(body)
}

// convertHTTPBodyToNote is used to convert user input on the frontend for use with the backend Notes list
func convertHTTPBodyToNote(httpBody io.ReadCloser) (notes.Note, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return notes.Note{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToNote(body)
}

func convertJSONBodyToTodo(jsonBody []byte) (todo.Todo, int, error) {
	var todoItem todo.Todo
	err := json.Unmarshal(jsonBody, &todoItem)
	if err != nil {
		return todo.Todo{}, http.StatusBadRequest, err
	}
	return todoItem, http.StatusOK, nil
}

func convertJSONBodyToNote(jsonBody []byte) (notes.Note, int, error) {
	var noteItem notes.Note
	err := json.Unmarshal(jsonBody, &noteItem)
	if err != nil {
		return notes.Note{}, http.StatusBadRequest, err
	}
	return noteItem, http.StatusOK, nil
}
