package todo_test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
)

func TestGet(t *testing.T) {
    got := todo.Get()
    if len(got) != 0 {
        t.Errorf("Array not empty! Size: %d", len(got))
    }
}

func TestAdd(t *testing.T) {
    // Don't allow empty task messages
    var index int = -1
    id := todo.Add("")
    list := todo.Get()
    for i, t := range list {
		if t.ID == id {
			index = i
		}
	}

    if index != -1 {
        t.Errorf("Added Task with empty message")
    }
}

// func TestDelete(t *testing.T) {
//     // Attempt to delete task that doesn't exist
//     got := todo.Delete("")
//     if len(got) != 0 {
//         t.Errorf("Array not empty! Size: %d", len(got))
//     }
// }

// func TestComplete(t *testing.T) {
//     // Attempt to complete task that doesn't exxist
//     got := todo.Complete("")
//     if len(got) != 0 {
//         t.Errorf("Array not empty! Size: %d", len(got))
//     }
// }