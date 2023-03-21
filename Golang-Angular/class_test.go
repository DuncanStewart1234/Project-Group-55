package todo_test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
)

func TestGet(t *testing.T) {
	list := course.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
}

func TestAdd(t *testing.T) {
	name := "Software Engineering"
	abv := "CEN3031"
	loc := "[-56.36,66.20]"
	str := `{"Mon":[["3:00","3:50"]],"Wed":[["3:00","3:50"]], "Fri":[["3:00","3:50"]]}`
	
    _, err := course.Add(name, abv, loc, str)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestDelete(t *testing.T) {
	err := course.Delete("372401")
	if err != nil {
		t.Errorf("error with delete function")
	}
}

// func Test_newPeriod(t *testing.T) {
// 	period := [["0","56"]]
// 	periods := course.newPeriod(period)
// }

// func Test_getLocationFromJSON(t *testing.T) {
// 	loc := "[0,56]"
// 	location := course.getLocationFromJSON(loc)
// }

// func Test_getScheduleFromJSON(t *testing.T) {
// 	sch := `{"Mon":[["3:00","3:50"]],"Wed":[["3:00","3:50"]], "Fri":[["3:00","3:50"]]}`
// 	schedule := course.getScheduleFromJSON(sch)
// }
