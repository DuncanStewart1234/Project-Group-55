// The notes package implements and maintains a list of notes entered by the user
package notes

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
	"github.com/rs/xid"
)

var (
	list []Note
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
	curr_uid int
)

// Note is a struct that holds info needed for notes list
type Note struct {
	ID      string `json:"id" gorm:"primarykey"`
	User_ID int    `json:"uid"`
	Title   string `json:"title" gorm:"size:256"`
	Category string `json:"category"`
	Message string `json:"message"`
}

// init is a constructor and calls initialiseList
func Start() {
	if user.GetUID() != 0 {
		once.Do(initialiseList)
	}
}

// initialiseList creates the notes list and calls initDatabase
func initialiseList() {
	list = []Note{}
	initDatabase()
}

// initDatabase initalises the database
func initDatabase() {
	db = utils.GetDB("src/server/databases/notes_list.db")
	db.AutoMigrate(&Note{})
}

func Close() {
	list = nil
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// Get returns the notes list
func Get() []Note {
	updateList()
	return list
}

// Add creates and adds a note to the notes list
func Add(title string, cat string, message string) (string, error) {
	mtx.Lock()
	updateList()
	err := utils.CheckIfEmptyOrTooLong(title)
	if err != nil {
		return "", err
	}

	t := newNote(title, cat, message)
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.ID, nil
}

// Edit finds a note in the list and edits its message
func Edit(id string, new_title string, new_cat string, new_msg string) error {
	updateList()
	location, err := findNoteLocation(id)
	if err != nil {
		return err
	}
	editNoteByLocation(location, new_title, new_cat, new_msg)
	db.Save(&list[location])
	return nil
}

// Delete removes and deletes a note from the notes list
func Delete(id string) error {
	updateList()
	location, err := findNoteLocation(id)
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newNote is a helper function to Add
func newNote(title string, cat string, msg string) Note {
	return Note {
		ID:      xid.New().String(),
		User_ID: curr_uid,
		Title:   title,
		Category: cat,
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
func editNoteByLocation(location int, title string, cat string, msg string) {
	mtx.Lock()
	if title != "" {
		list[location].Title = title
	}

	if cat != "" {
		list[location].Category = cat
	}

	if msg != "" {
		list[location].Message = msg
	}
	mtx.Unlock()
}

func updateList() error {
	curr_uid = user.GetUID()
	result := db.Where("User_ID = ?", curr_uid).Find(&list)
	return result.Error
}