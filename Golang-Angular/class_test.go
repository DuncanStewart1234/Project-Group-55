package test

import (
	"fmt"
	"testing"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
)

func TestCourseGet(t *testing.T) {
	list := course.Get()
	fmt.Println(list)
	if list == nil {
		t.Errorf("error with get function")
	}
}

func TestCourseAdd(t *testing.T) {
	name := "Programming Language Concepts"
	abv := "COP4020"
	loc := "[-36.36,66.20]"
	str := `{"Mon":[["12:50","13:40"]],"Wed":[["12:50","13:40"]], "Fri":[["12:50","13:40"]]}`
	
    _, err := course.Add(name, abv, loc, str)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestCourseCalAdd(t *testing.T) {
	name := "Computer Aided Animation"
	loc := `{"lat": 38.641546, "lng": -45.351659}`
	start := "2023-04-17T8:30:00"
	end := "2023-04-17T9:20:00"
	
    _, err := course.AddCal(name, loc, start, end)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestCourseDelete(t *testing.T) {
	err := course.Delete("283917")
	if err != nil {
		t.Errorf("error with delete function")
	}
}

func TestCourseEdit(t *testing.T) {
	// name := "Software Engineering 2"
	// abv := "CEN3032"
	// loc := "[-566.36,665.20]"
	name := ""
	abv := ""
	loc := ""
	// str := `{"Mon":[["3:00","3:50"]],"Wed":[["3:00","3:50"]]}`
	start := "2023-04-17T7:30:00"
	end := "2023-04-17T9:20:00"

	err := course.Edit(201953, name, abv, loc, start, end)
	if err != nil {
		t.Errorf("error with delete function")
	}
}
