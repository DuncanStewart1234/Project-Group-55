package test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/schedule"
)

func TestScheduleGet(t *testing.T) {
	list := schedule.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
}

func TestScheduleAdd(t *testing.T) {
	cid := "211479"
    _, err := schedule.Add(cid)
    if err != nil {
		t.Errorf("error with add function")
	}
}

func TestScheduleDelete(t *testing.T) {
	err := schedule.Delete("cgdk46siuh41bo3t6100")
	if err != nil {
		t.Errorf("error with delete function")
	}
}

func TestScheduleGetUser(t *testing.T) {
	uid := 1005
	list := schedule.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
	
	for _, sch := range list {
		if sch.User_ID != uid {
			t.Errorf("returned schedule that isn't the current users")
		}
	}

}