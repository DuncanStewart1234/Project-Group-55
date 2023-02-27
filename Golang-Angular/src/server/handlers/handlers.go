package handlers

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"

    "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
    "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
    "github.com/gin-gonic/gin"
)


// To Do Handlers

func GetTodoListHandler(c *gin.Context) {
    c.JSON(http.StatusOK, todo.Get())
}

func AddTodoHandler(c *gin.Context) {
    todoItem, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
    if err != nil {
        c.JSON(statusCode, err)
        return
    }

    id, _ := todo.Add(todoItem.Message)
    c.JSON(statusCode, gin.H{"id": id})
}

func DeleteTodoHandler(c *gin.Context) {
    todoID := c.Param("id")
    if err := todo.Delete(todoID); err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, "")
}

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

func GetNotesHandler(c *gin.Context) {
    c.JSON(http.StatusOK, notes.Get())
}

func AddNotesHandler(c *gin.Context) {
    notesItem, statusCode, err := convertHTTPBodyToNote(c.Request.Body)
    if err != nil {
        c.JSON(statusCode, err)
        return
    }
    c.JSON(statusCode, gin.H{"id": notes.Add(notesItem.Title, notesItem.Message)})
}

func DeleteNotesHandler(c *gin.Context) {
    noteID := c.Param("id")
    if err := notes.Delete(noteID); err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, "peration Completed Successfully!")
}

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


func convertHTTPBodyToTodo(httpBody io.ReadCloser) (todo.Todo, int, error) {
    body, err := ioutil.ReadAll(httpBody)
    if err != nil {
        return todo.Todo{}, http.StatusInternalServerError, err
    }
    defer httpBody.Close()
    return convertJSONBodyToTodo(body)
}

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