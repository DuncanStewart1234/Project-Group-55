# Backend API Routes

## User Account API
`GET("/account/:user")`
- Get user account information

`POST("/signup")`
- Sign up for user account

`POST("/login")`
- Log into user account

`DELETE("/:uid/delete")`
- Delete user account

`POST("/logout")`
- Log out of current account

`PUT("/account/edit")`
- Edit user account details

---
## Weather API
`GET("/weather/forecast")`
- Returns 

`GET("/weather")`
- Returns 

---
## Schedules REST API
`GET("/schedule")`
- Returns user schedule

`POST("/schedule")`
- Adds class to user schedule

`DELETE("/schedule/:id")`
- Removes class from user schedule

---
## Course REST API
`GET("/course")`
- Returns all classes

`POST("/course")`
- Adds new class

`DELETE("/course/:id")`
- Deletes class

---
## Notes REST API
`GET("/notes")`
- Gets all notes for current user

`POST("/notes")`
- Adds note for current user

`DELETE("/notes/:id")`
- Deletes note

`PUT("/notes")`
- Edits note

---
## Todo REST API
`GET("/todo")`
- Gets all tasks for current user

`POST("/todo")`
- Adds task for current user

`DELETE("/todo/:id")`
- Removes task

`PUT("/todo")`
- Sets task as complete
