// This package implements and maintains handler functions for our application
package handlers

import (
	"encoding/json"
	"io"
	// TODO: Update
	"io/ioutil"
	"net/http"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/weather"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils/token"

	"github.com/gin-gonic/gin"
)

// To Do Handlers
// GetTodoListHandler is a handler to request the todo list from the Todo package
func GetTodoListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todo.Get())
}

// AddTodoHandler is used to add to the todo list from the Todo package
func AddTodoHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}

	id, _ := todo.Add(item.Message)
	c.JSON(statusCode, gin.H{"id": id})
}

// DeleteTodoHandler is used to delete an item from the todo list
func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id")
	if err := todo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "task deleted successfully")
}

// CompleteTodoHandler is used to mark a todo list item as complete
func CompleteTodoHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToTodo(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if todo.Complete(item.ID) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "task completed successfully")
}



// Notes Handlers
// GetNotesHandler is used to request the notes list from the Notes package
func GetNotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, notes.Get())
}

// AddNotesHandler is used to add a note to the notes list from the Notes package
func AddNotesHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToNote(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	id, _ := notes.Add(item.Title, item.Message)
	c.JSON(statusCode, gin.H{"id": id})
}

// DeleteNotesHandler is used to delete a note from the notes list
func DeleteNotesHandler(c *gin.Context) {
	id := c.Param("id")
	if err := notes.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "note deleted successfully")
}

// EditNotesHandler is used to edit a note from the notes list
func EditNotesHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToNote(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	if notes.Edit(item.ID, item.Title, item.Message) != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "note edited successfully")
}



// Schedule + Maps Handlers
func GetSchedulesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, schedule.Get())
}

func AddSchedulesHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToSchedule(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}

	id, _ := schedule.Add(item.Class_ID)
	c.JSON(statusCode, gin.H{"id": id})
}

func DeleteSchedulesHandler(c *gin.Context) {
	id := c.Param("id")
	if err := schedule.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "schedule deleted successfully")
}


// Admin Routes Handlers
func GetAdminUsersHandler (c *gin.Context) {
	// uname := c.Param("user")
	err := token.TokenValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}	

	c.JSON(http.StatusOK, user.GetAll())
}

func DeleteAdminUsersHandler(c *gin.Context) {
	id := c.Param("uid")
	if err := user.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "user deleted successfully")
}


// Users Handlers
func GetUserHandler (c *gin.Context) {
	uname := c.Param("user")
	err := token.TokenValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}	

	c.JSON(http.StatusOK, gin.H{"User": user.Get(uname)})
}

func SignupHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToUserSignup(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}	

	uid, _ := user.Add(item.First_Name, item.Last_Name, item.User_Name, item.Email, item.Password)
	c.JSON(statusCode, gin.H{"uid": uid})
}

func LoginHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToUserLogin(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}	
	
	token, err := user.Login(item.Username, item.Password)
	if err != nil {
		c.JSON(400, err)
		return
	}	

	c.JSON(statusCode, gin.H{"token": token})
}

func DeleteUsersHandler(c *gin.Context) {
	id := c.Param("uid")
	if err := user.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "user deleted successfully")
}



// Class Handlers
func GetClassesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, course.Get())
}

func AddClassHandler(c *gin.Context) {
	item, statusCode, err := convertHTTPBodyToClass(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}

	// cid, _ := course.Add(item.Name, item.Abbrv, item.Location, item.Schedule)
	cid, _ := course.AddCal(item.Title, item.ExtendedProps, item.Start, item.End)
	c.JSON(statusCode, gin.H{"cid": cid})
}

func DeleteClassHandler(c *gin.Context) {
	id := c.Param("cid")
	if err := course.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Operation Completed Successfully!")
}

// func EditClassHandler(c *gin.Context) {
// 	item, statusCode, err := convertHTTPBodyToClass(c.Request.Body)
// 	if err != nil {
// 		c.JSON(statusCode, err)
// 		return
// 	}
// 	if course.Edit(item.Class_ID, item.Name, item.Abbrv, item.Location, item.Schedule) != nil {
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, "Operation Completed Successfully!")
// }



// Weather Handlers
func GetWeatherHandler(c *gin.Context) {
	c.JSON(http.StatusOK, weather.GetWeather("Gainesville", "f", "en"))
}

func GetWeatherForecastHandler(c *gin.Context) {
	c.JSON(http.StatusOK, weather.GetForecast("Gainesville", "f", "en"))
}



// HTTP To Class Objects
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

func convertHTTPBodyToClass(httpBody io.ReadCloser) (CalClass, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return CalClass{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToClass(body)
}

func convertHTTPBodyToSchedule(httpBody io.ReadCloser) (schedule.StudentSchedule, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return schedule.StudentSchedule{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToSchedule(body)
}

func convertHTTPBodyToUserSignup(httpBody io.ReadCloser) (RegisterInput, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return RegisterInput{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToUserSignup(body)
}

func convertHTTPBodyToUserLogin(httpBody io.ReadCloser) (LoginInput, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return LoginInput{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToUserLogin(body)
}



// JSON to Struct Objects
func convertJSONBodyToTodo(jsonBody []byte) (todo.Todo, int, error) {
	var item todo.Todo
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return todo.Todo{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}

func convertJSONBodyToNote(jsonBody []byte) (notes.Note, int, error) {
	var item notes.Note
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return notes.Note{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}

func convertJSONBodyToClass(jsonBody []byte) (CalClass, int, error) {
	var item CalClass
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return CalClass{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}

func convertJSONBodyToSchedule(jsonBody []byte) (schedule.StudentSchedule, int, error) {
	var item schedule.StudentSchedule
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return schedule.StudentSchedule{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}

func convertJSONBodyToUserSignup(jsonBody []byte) (RegisterInput, int, error) {
	var item RegisterInput
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return RegisterInput{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}

func convertJSONBodyToUserLogin(jsonBody []byte) (LoginInput, int, error) {
	var item LoginInput
	err := json.Unmarshal(jsonBody, &item)
	if err != nil {
		return LoginInput{}, http.StatusBadRequest, err
	}
	return item, http.StatusOK, nil
}


func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}


// Struct Classes
type HandlerClass struct {
	ID uint
	Class_ID int `json:"cid"`
	Name string `json:"name"`
	Abbrv string `json:"abbrv"`
	Location string `json:"loc"`
	Schedule string `json:"schedule"`
}

type CalClass struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End string `json:"end"`
	ExtendedProps string `json:"extendedProps"`
}

type RegisterInput struct {
	First_Name string `json:"first"`
	Last_Name string `json:"last"`
	Email string `json:"email"`
	User_Name string `json:"uname"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}