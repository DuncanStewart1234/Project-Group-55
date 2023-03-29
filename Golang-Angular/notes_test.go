package test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/notes"

)

type addTestNote struct {
	arg string
}

var (
	testList     []notes.Note
	addTestNotes = []addTestNote{
		{"Message"},
		{"Hello"},
		{"testing testing"},
		{"test4"},
		{"final test"},
	}
)

func TestAddNote(t *testing.T) {

	for _, test := range addTestNotes {
		notes.Add("Title", test.arg)
		want := false

		testList = notes.Get()

		for i, t := range testList {
			if t.Message == test.arg {
				want = true
				_ = i
				continue
			}
		}

		if want == false {
			t.Errorf("note was not added")
		}
	}
}

func TestNoteGetUser(t *testing.T) {
	uid := 1005
	list := notes.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
	
	for _, sch := range list {
		if sch.User_ID != uid {
			t.Errorf("returned schedule that isn't the current users")
		}
	}

}