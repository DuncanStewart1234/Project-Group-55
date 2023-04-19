package databases

import (
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
)

// TODO: Move to Login Function
func DB_Online() {
	if user.GetUID() != 0 {
		course.Start()
		notes.Start()
		schedule.Start()
		todo.Start()
	}
}

func DB_Offline() {
	if user.GetUID() != 0 {
		course.Close()
		notes.Close()
		schedule.Close()
		todo.Close()
	}
}
