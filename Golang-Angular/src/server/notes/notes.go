package notes

import (
	"errors"
	"sync"
    "gorm.io/gorm"

    "github.com/rs/xid"
    "github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []Note
    db *gorm.DB
	mtx  sync.RWMutex
    once sync.Once
)

type Note struct {
    gorm.Model
	ID		string `json:"id"`
    User_ID string `json:"uid"`
	Title	string `json:"title"`
	Message	string `json:"message"`
}

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Note{}
    initDatabase()
}

func initDatabase() {
    db = utils.GetDB("src/server/databases/notes_list.db")

    db.AutoMigrate(&Note{})

    result := db.Find(&list)
    if result.Error	!= nil {
        panic("failed to connect database")
    }
}

func Get() []Note {
    return list
}

func Add(title string, message string) string {
    t := newNote(title, message)
    mtx.Lock()
    list = append(list, t)
    db.Create(&t);
    mtx.Unlock()
    return t.ID
}

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

func Delete(id string) error {
    location, err := findNoteLocation(id)
    if err != nil {
        return err
    }
    db.Delete(&list[location])
    removeElementByLocation(location)
    return nil
}

func newNote(title string, msg string) Note {
    // TODO: Get UID
    return Note {
        ID:		xid.New().String(),
        User_ID: "1005",
        Title: 	title,
        Message: msg,
    }
}

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

func removeElementByLocation(i int) {
    mtx.Lock()
    list = append(list[:i], list[i+1:]...)
    mtx.Unlock()
}

func editNoteByLocation(location int, msg string) {
    mtx.Lock()
    list[location].Message = msg
    mtx.Unlock()
}
