package test

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
        t.Errorf("Failed to delete task")
    }
}

func TestEmptyMessageAdd(t *testing.T) {
    // Don't Allow Empty Task Message
    _, err := todo.Add("")
    if err == nil {
        t.Errorf("Added Task with empty message")
    }
}

func TestLargeMessageAdd(t *testing.T) {
    // Don't Allow Empty Task Message
    _, err := todo.Add("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?")
    if err == nil {
        t.Errorf("Added Task with large message")
    }
}

func TestDeleteImaginaryTask(t *testing.T) {
    // Attempt To Delete Task That Does Not Exist
    err := todo.Delete("")
    if err == nil {
        t.Errorf("Failed to capture error")
    }
}

func TestTodoGetUser(t *testing.T) {
	uid := 1005
	list := todo.Get()
	if list == nil {
		t.Errorf("error with get function")
	}
	
	for _, task := range list {
		if task.User_ID != uid {
			t.Errorf("returned task that isn't the current users")
		}
	}

}