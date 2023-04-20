// This package implements and maintains a todo list for the user to work with.
package todo

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/user"
	"github.com/rs/xid"
)

var (
	list []Todo
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
	curr_uid int
)

// Todo is a struct model containing the Todo list item info
type Todo struct {
	ID       string `json:"id" gorm:"primarykey"`
	User_ID  int	`json:"uid"`
	Message  string `json:"message" gorm:"size:256"`
	Complete bool   `json:"complete"`
}

// Start is a constructor, calls the initialise list function
func Start() {
	if user.GetUID() != 0 {
		once.Do(initialiseList)
	}
}

// initialiseList initialises the Todo list
func initialiseList() {
	list = []Todo{}
	initDatabase()
}

// initDatabase initalises the database
func initDatabase() {
	db = utils.GetDB("src/server/databases/todo_list.db")
	db.AutoMigrate(&Todo{})
}

// Close closes the SQL database
func Close() {
	list = nil
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// Get returns the Todo list
func Get() []Todo {
	updateList()
	return list
}

// Add a new item to the Todo list based off message input
func Add(message string) (string, error) {
	mtx.Lock()
	updateList()
	err := utils.CheckIfEmptyOrTooLong(message)
	if err != nil {
		return "", err
	}

	t := newTodo(message)
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.ID, nil
}

// Complete marks a Todo list item as complete
func Complete(id string) error {
	updateList()
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	setTodoCompleteByLocation(location)
	db.Save(&list[location])
	return nil
}

// Delete a Todo list item
func Delete(id string) error {
	updateList()
	location, err := findTodoLocation(id)
	if err != nil {
		return err
	}
	db.Where("ID = ?", list[location].ID).Unscoped().Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newTodo creates a new Todo item, helper function to Add
func newTodo(msg string) Todo {
	return Todo {
		ID:       xid.New().String(),
		User_ID:  curr_uid,
		Message:  msg,
		Complete: false,
	}
}

// findTodoLocation searches the list for an item
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

// removeElementByLocation removes an item
func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}

// setTodoCompleteByLocation is a helper function to Complete
func setTodoCompleteByLocation(location int) {
	mtx.Lock()
	list[location].Complete = true
	mtx.Unlock()
}

func updateList() error {
	curr_uid = user.GetUID()
	result := db.Where("User_ID = ?", curr_uid).Find(&list)
	return result.Error
}
