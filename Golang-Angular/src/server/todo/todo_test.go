package todo

import (
	"testing"
)

type addTodoItem struct {
	arg string
}

var (
	testList    []Todo
	addTestTodo = []addTodoItem{
		{"Message"},
		{"Hello"},
		{"testing testing"},
		{"test4"},
		{"final test"},
	}
)

func TestAddTodoItem(t *testing.T) {

	for _, test := range addTestTodo {
		Add(test.arg)
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
