package todo

import (
	"errors"
	"sync"
    "gorm.io/gorm"

    "github.com/rs/xid"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []Todo
	db *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

type Todo struct {
	gorm.Model
	ID       string `json:"id"`
	User_ID	 string `json:"uid"`
	Message  string `json:"message" gorm:"size:256"`
	Complete bool   `json:"complete"`
}

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Todo{}
	initDatabase()
}

func initDatabase() {
	db = utils.GetDB("src/server/databases/todo_list.db")

    db.AutoMigrate(&Todo{})

	result := db.Find(&list)
    if result.Error	!= nil {
        panic("failed to connect database")
    }
}

func Get() []Todo {
	return list
}

func Add(message string) (string, error) {
	t := newTodo(message)
	if message == "" {
		return "", errors.New("message cannot be empty")
	}
	mtx.Lock()
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.ID, nil
}

func Complete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	db.Save(&list[location])
	return nil
}

func Delete(id string) error {
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	db.Where("ID = ?", list[location].ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

func newTodo(msg string) Todo {
	// TODO: Get UID
	return Todo {
		ID:       xid.New().String(),
		User_ID: "1005",
		Message:  msg,
		Complete: false,
	}
}

func findTodoLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if t.ID == id {
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
