package notes

import (
	"errors"
	"sync"

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
    return t.ID
}

func Delete(id string) error {
    location, err := findNoteLocation(id)
    if err != nil {
        return err
    }
    removeElementByLocation(location)
    return nil
}

func Edit(id string, new_msg string) error {
    location, err := findNoteLocation(id)
    if err != nil {
        return err
    }
    EditNoteByLocation(location, new_msg)
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

