package notes

import (
	"testing"
)

type addTestNote struct {
	arg string
}

var (
	testList     []Note
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
		Add("Title", "", test.arg)
		want := false

		testList = Get()

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
