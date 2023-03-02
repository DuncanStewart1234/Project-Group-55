// The notes package implements and maintains a list of notes entered by the user
package notes

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/rs/xid"
)

var (
	list []Note
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

// Note is a struct that holds info needed for notes list
type Note struct {
	gorm.Model
	ID      string `json:"id"`
	User_ID string `json:"uid"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// init is a constructor and calls initialiseList
func init() {
	once.Do(initialiseList)
}

// initialiseList creates the notes list and calls initDatabase
func initialiseList() {
	list = []Note{}
	initDatabase()
}

// initDatabase initalises the database
func initDatabase() {
	db = utils.GetDB("C:/Users/scott/go/src/github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/databases/notes_list.db")

	db.AutoMigrate(&Note{})

	result := db.Find(&list)
	if result.Error != nil {
		panic("failed to connect database")
	}
}

// Get returns the notes list
func Get() []Note {
	return list
}

// Add creates and adds a note to the notes list
func Add(title string, message string) string {
	t := newNote(title, message)
	mtx.Lock()
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.ID
}

// Edit finds a note in the list and edits its message
func Edit(id string, new_msg string) error {
	// TODO: Allow for title edit
	location, err := findNoteLocation(id)
	if err != nil {
		return err
	}
	editNoteByLocation(location, new_msg)
	db.Save(&list[location])
	return nil
}

// Delete removes and deletes a note from the notes list
func Delete(id string) error {
	location, err := findNoteLocation(id)
	if err != nil {
		return err
	}
	db.Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newNote is a helper function to Add
func newNote(title string, msg string) Note {
	// TODO: Get UID
	return Note{
		ID:      xid.New().String(),
		User_ID: "1005",
		Title:   title,
		Message: msg,
	}
}

// findNoteLocation is a helper function to Edit and Delete
func findNoteLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if t.ID == id {
			return i, nil
		}
	}
	return 0, errors.New("could not find note based on id")
}

// removeElementByLocation is a helper function to Delete
func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

// editNoteByLocation is a helper function to Edit
func editNoteByLocation(location int, msg string) {
	mtx.Lock()
	list[location].Message = msg
	mtx.Unlock()
}
