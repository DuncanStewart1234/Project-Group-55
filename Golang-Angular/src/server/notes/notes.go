package notes

import (
    "database/sql"
	"errors"
	"sync"
    "strconv"

	_ "github.com/mattn/go-sqlite3"
    "github.com/rs/xid"
)

var (
	list []Note
	mtx  sync.RWMutex
    once sync.Once
)

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Note{}
    initDatabase()
}

type Note struct {
	ID		string `json:"id"`
	Title	string `json:"title"`
	Message	string `json:"message"`
}

func Get() []Note {
    return list
}

func Add(title string, message string) string {
    t := newNote(title, message)
    mtx.Lock()
    list = append(list, t)
    mtx.Unlock()
    // Example UserID 1005
    insertDBEntry(t.ID, 1005, title, message)
    return t.ID
}

func Delete(id string) error {
    location, err := findNoteLocation(id)
    if err != nil {
        return err
    }
    removeElementByLocation(location)
    deleteDBEntry(id)
    return nil
}

func Edit(id string, new_msg string) error {
    location, err := findNoteLocation(id)
    if err != nil {
        return err
    }
    EditNoteByLocation(location, new_msg)
    editDBEntry(id, new_msg)
    return nil
}

func newNote(title string, msg string) Note {
    return Note{
        ID:		xid.New().String(),
        Title: 	title,
        Message: msg,
    }
}

func findNoteLocation(id string) (int, error) {
    mtx.RLock()
    defer mtx.RUnlock()
    for i, t := range list {
        if isMatchingID(t.ID, id) {
            return i, nil
        }
    }
    return 0, errors.New("could not find note based on id")
}

func removeElementByLocation(i int) {
    mtx.Lock()
    list = append(list[:i], list[i+1:]...)
    mtx.Unlock()
}

func EditNoteByLocation(location int, msg string) {
    mtx.Lock()
    list[location].Message = msg
    mtx.Unlock()
}

func isMatchingID(a string, b string) bool {
    return a == b
}

func initDatabase() {
    db, _ := sql.Open("sqlite3", "src/server/databases/notes_list.db")
    create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS notes (id TEXT PRIMARY KEY, user_id INTEGER NOT NULL, title TEXT NOT NULL, message TEXT NOT NULL)")
    create.Exec()

	rows, _ := db.Query("SELECT id, title, message FROM notes")

	var id string
	var title string
	var message string

    for rows.Next() {
        rows.Scan(&id, &title, &message)
        t := Note {
            ID:       id,
            Title: title,
            Message:  message,
        }
        list = append(list, t)
    }

	rows.Close()
}

func deleteDBEntry(id string) {
	db, _ := sql.Open("sqlite3", "src/server/databases/notes_list.db")
	statement, _ := db.Prepare("DELETE FROM notes WHERE id=?")
	statement.Exec(id)
}

func editDBEntry(id string, msg string) {
	db, _ := sql.Open("sqlite3", "src/server/databases/notes_list.db")
	statement, _ := db.Prepare("UPDATE notes SET message=? WHERE id=?")
	statement.Exec(msg, id)
}

func insertDBEntry(id string, user_id int, title string, message string) {
    db, _ := sql.Open("sqlite3", "src/server/databases/notes_list.db")
    statement, _ := db.Prepare("INSERT INTO notes VALUES (?, ?, ?, ?)")
    statement.Exec(id, strconv.Itoa(user_id), title, message)
}