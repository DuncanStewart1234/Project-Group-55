package test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
)

func TestCourseGet(t *testing.T) {
	list := course.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
}

func TestCourseAdd(t *testing.T) {
	name := "Software Engineering"
	abv := "CEN3031"
	loc := "[-56.36,66.20]"
	str := `{"Mon":[["3:00","3:50"]],"Wed":[["3:00","3:50"]], "Fri":[["3:00","3:50"]]}`
	
    _, err := course.Add(name, abv, loc, str)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestCourseDelete(t *testing.T) {
	err := course.Delete("974174")
	if err != nil {
		t.Errorf("error with delete function")
	}
}

func TestCourseEdit(t *testing.T) {
	name := "Software Engineering 2"
	abv := "CEN3032"
	loc := "[-566.36,665.20]"
	str := `{"Mon":[["3:00","3:50"]],"Wed":[["3:00","3:50"]]}`

	err := course.Edit("215679", name, abv, loc, str)
	if err != nil {
		t.Errorf("error with delete function")
	}
}
