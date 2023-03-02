// A package used to create a schedule of classes that a college student
// may use.
package schedule

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/rs/xid"
)

var (
	list []Class
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

// Class is a struct used to contain info about a student's class
type Class struct {
	gorm.Model
	// ID       string `json:"id"`
	Class_ID string `json:"cid"`
	Location string `json:"loc"`
	// Add Schedule
}

// init is a constructor that calls initialiseList
func init() {
	once.Do(initialiseList)
}

// initialiseList creates an array called Class and calls initDatabase
func initialiseList() {
	list = []Class{}
	initDatabase()
}

// initDatabase initialises the databse for this package
func initDatabase() {
	db = utils.GetDB("C:/Users/scott/go/src/github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/databases/classes.db")

	db.AutoMigrate(&Class{})

	result := db.Find(&list)
	if result.Error != nil {
		panic("failed to connect database")
	}
}

// Get returns the list of classes in the schedule
func Get() []Class {
	return list
}

// Add creates and adds a class element to the list
func Add(loc string) (string, error) {
	t := newClass(loc)
	// if loc == "" {
	// return "", errors.New("lo cannot be empty")
	// }
	mtx.Lock()
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.Class_ID, nil
}

// Delete removes and deletes a class element from the list
func Delete(cid string) error {
	location, err := findClassLocation(cid)
	if err != nil {
		return err
	}
	db.Where("Class_ID = ?", list[location].Class_ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newClass is a helper function to Add
func newClass(loc string) Class {
	return Class{
		Class_ID: xid.New().String(),
		Location: loc,
	}
}

// findClassLocation is a helper function to Delete
func findClassLocation(cid string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if t.Class_ID == cid {
			return i, nil
		}
	}
	return 0, errors.New("could not find class based on cid")
}

// removeElementByLocation is a helper function to Delete
func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}
