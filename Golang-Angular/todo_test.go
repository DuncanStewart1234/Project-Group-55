package todo_test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/todo"
)

func TestDB(t *testing.T) {
    id, err := todo.Add("Test Todo Task")
    if err != nil {
		t.Errorf("Add error")
	}
    list := todo.Get()
    for _, task := range list {
        if id == task.ID {
            if task.Message != "Test Todo Task" {
                t.Errorf("Did not add task")
            }
            break
        }
    }

    err = todo.Delete(id)
    if err != nil {
        t.Errorf("Failed to capture error")
    }
}

func TestEmptyMessageAdd(t *testing.T) {
    // Don't allow empty task messages
    _, err := todo.Add("")
    if err == nil {
        t.Errorf("Added Task with empty message")
    }
}

func TestDeleteImaginaryTask(t *testing.T) {
    // Attempt to delete task that doesn't exist
    err := todo.Delete("")
    if err == nil {
        t.Errorf("Failed to capture error")
    }
}

// func TestComplete(t *testing.T) {
//     // Attempt to complete task that doesn't exxist
//     got := todo.Complete("")
//     if len(got) != 0 {
//         t.Errorf("Array not empty! Size: %d", len(got))
//     }
// }