package test

import (
	"fmt"
	"strconv"
	"testing"
	
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"


)


func TestProgram(t *testing.T) {
	// TODO: Randomize and loop
	fname := "Endrick"
	lname := "Lafosse"
	uname := "elafosse"
	email := "elafosse@gmail.com"
	passwd := "password"

	// Creates User
    _, err := user.Add(fname, lname, uname, email, passwd)
    if err != nil {
		t.Errorf("error with add function")
	}

	
	// Logs User In
	token, err := user.Login(uname, passwd)
	fmt.Println(token)
	if err != nil || token == "" {
		t.Errorf("error with login function")
	}
	
	// Gets User Account Details
	data := user.Get(uname)
	fmt.Println(data)
	if data.First_Name == "" {
		t.Errorf("error with get function")
	}
	
	// Edits User
	user.Edit(strconv.Itoa(data.User_ID), "Kendrick", "Lamar", "kungfukenny", "kdot@gmail.com", "", passwd)
	
	// Testing Todo
    task_id_one, _ := todo.Add("Finish Project")
    task_id_two, _ := todo.Add("Work on Album")

	todo.Complete(task_id_one)
	todo.Complete(task_id_two)

    tasks := todo.Get()
	fmt.Println(tasks)
	for _, task := range tasks {
		if task.User_ID != data.User_ID {
			t.Errorf("returned task that isn't the current users")
		}
	}

    todo.Delete(task_id_one)
    todo.Delete(task_id_two)

	
	// Testing Notes
	note_id_one, _ := notes.Add("Alrights", "So I Went Running For Answers")
	note_id_two, _ := notes.Add("Money Trees", "Dreams of me getting shaded under a money tree")
	notes.Edit(note_id_one, "Alright", "")
	notes.Edit(note_id_two, "", "Dreams of Me Getting Shaded Under A Money Tree")

	notes_list := notes.Get()
	fmt.Println(notes_list)
	for _, note := range notes_list {
		if note.User_ID != data.User_ID {
			t.Errorf("returned note that isn't the current users")
		}
	}
	
	notes.Delete(note_id_one)
	notes.Delete(note_id_two)


	// Testing Classes & Schedule
	class_one_name := "Programming Language Concepts"
	class_one_abv := "COP4020"
	class_one_loc := "[-36.36,66.20]"
	class_one_sch := `{"Mon":[["12:50","13:40"]],"Wed":[["12:50","13:40"]], "Fri":[["12:50","13:40"]]}`

	class_two_name := "Intro to Computer Aided Animation"
	class_two_abv := "CAP3034"
	class_two_loc := `{"lat": 38.641546, "lng": -45.351659}`
	class_two_start := "2023-04-17T8:30:00"
	class_two_end := "2023-04-17T9:20:00"
	
    class_one_id, _ := course.Add(class_one_name, class_one_abv, class_one_loc, class_one_sch)
    class_two_id, _ := course.AddCal(class_two_name, class_two_abv, class_two_loc, class_two_start, class_two_end)

	schedule_one_id, _ := schedule.Add(strconv.Itoa(class_one_id))
	schedule_two_id, _ := schedule.Add(strconv.Itoa(class_two_id))

	classes := course.Get()
	user_schedule := schedule.Get()
	fmt.Println(classes)
	fmt.Println(user_schedule)
	for _, sch := range user_schedule {
		if sch.User_ID != data.User_ID {
			t.Errorf("returned schedule that isn't the current users")
		}
	}

	schedule.Delete(schedule_one_id)
	schedule.Delete(schedule_two_id)
	course.Delete(strconv.Itoa(class_one_id))
	course.Delete(strconv.Itoa(class_two_id))


	// Deletes User Account
	user.Delete(strconv.Itoa(data.User_ID))
	// TODO: Check if deleted

	// Logs User Out
	// user.Logout()
	// TODO: Check if logged out
}
