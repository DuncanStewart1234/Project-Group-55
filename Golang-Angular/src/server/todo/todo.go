package todo

import (
	"errors"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
)

var (
	list []Todo
	mtx  sync.RWMutex
	once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Todo{}
	/////initDatabase()
}

// Todo data structure for a task with a description of what to do
type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

// Get retrieves all elements from the todo list
func Get() []Todo {
	return list
}

// Add will add a new todo based on a message
func Add(message string) string {
	t := newTodo(message)
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	/////insertDBEntry(t.ID, 1005, message)
	return t.ID
}

// Delete will remove a Todo from the Todo list
func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	removeElementByLocation(location)
	//////deleteDBEntry(id)
	return nil
}

// Complete will set the complete boolean to true, marking a todo as
// completed
func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	/////completeDBEntry(id)
	return nil
}

func newTodo(msg string) Todo {
	return Todo{
		ID:       xid.New().String(),
		Message:  msg,
		Complete: false,
	}
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if isMatchingID(t.ID, id) {
			return i, nil
		}
	}
	return 0, errors.New("could not find todo based on id")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
	return a == b
}

/*func initDatabase() {
    db, _ := sql.Open("sqlite3", "src/server/databases/todo_list.db")
    create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS todo (id TEXT PRIMARY KEY, user_id INTEGER NOT NULL, message TEXT NOT NULL, is_complete bool NOT NULL)")
    create.Exec()

    rows, _ := db.Query("SELECT id, message, is_complete FROM todo")

    var id string
    var message string
    var is_complete bool

    for rows.Next() {
        rows.Scan(&id, &message, &is_complete)
        Add(message)
        if is_complete {
            Complete(id)
        }
    }

    rows.Close()
}

func deleteDBEntry(id string) {
    db, _ := sql.Open("sqlite3", "src/server/databases/todo_list.db")
    statement, _ := db.Prepare("DELETE FROM todo WHERE id=?")
    statement.Exec(id)
}

func completeDBEntry(id string) {
    db, _ := sql.Open("sqlite3", "src/server/databases/todo_list.db")
    statement, _ := db.Prepare("UPDATE todo SET is_complete=true WHERE id=?")
    statement.Exec(id)
}

func insertDBEntry(id string, user_id int, message string) {
    db, _ := sql.Open("sqlite3", "src/server/databases/todo_list.db")
    statement, _ := db.Prepare("INSERT INTO todo VALUES (?, ?, ?, ?)")
    statement.Exec(id, strconv.Itoa(user_id), message , false)
    statement.Exec()
}*/
