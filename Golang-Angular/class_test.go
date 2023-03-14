package todo_test

import (
	"testing"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/course"
)

func TestGet(t *testing.T) {
	list := course.Get()
	if list == nil {
		t.Errorf("Get error")
	}
}


func TestAdd(t *testing.T) {
	name := "Software Engineering 2"
	ab := "CEN3032"
	lat := "[-56.36,66.20]"
	str := `{"Mon":[["3:00","3:50"]],"Wed":[["11:30","12:20"],["3:00","3:50"]]}`
	
    _, err := course.Add(name, ab, lat, str)
    if err != nil {
		t.Errorf("Add error")
	}
}

func TestDelete(t *testing.T) {
	a := course.Delete("cg33q1ciuh43l64c9l7g")
	if a != nil {
		t.Errorf("Add error")
	}
}
// func TestEmptyMessageAdd(t *testing.T) {
//     // Don't Allow Empty Task Message
//     _, err := todo.Add("")
//     if err == nil {
//         t.Errorf("Added Task with empty message")
//     }
// }

// func TestLargeMessageAdd(t *testing.T) {
//     // Don't Allow Empty Task Message
//     _, err := todo.Add("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?")
//     if err == nil {
//         t.Errorf("Added Task with large message")
//     }
// }

// func TestDeleteImaginaryTask(t *testing.T) {
//     // Attempt To Delete Task That Does Not Exist
//     err := todo.Delete("")
//     if err == nil {
//         t.Errorf("Failed to capture error")
//     }
// }
