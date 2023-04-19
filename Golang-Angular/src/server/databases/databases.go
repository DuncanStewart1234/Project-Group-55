package databases

import (
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
)

// TODO: Move to Login Function
func DB_Online() {
	// course.Start()
	notes.Start()
	schedule.Start()
	todo.Start()
}

func DB_Offline() {
	course.Close()
	notes.Close()
	schedule.Close()
	todo.Close()
}
