// The schedule package implements and maintains a schedule of classes for a student
package schedule

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/rs/xid"
)

var (
	list []StudentSchedule
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

// StudentSchedule is the struct used to create a class schedule
type StudentSchedule struct {
	gorm.Model
	ID       string `json:"id"`
	User_ID  string `json:"uid"`
	Class_ID string `json:"cid"`
}


// init is a constructor, calls initialiseList
func init() {
	once.Do(initialiseList)
}

// initialiseList creates the schedule array and calls initDatabase
func initialiseList() {
	list = []StudentSchedule{}
	initDatabase()
}

// initDatabase initialises the database
func initDatabase() {
	db = utils.GetDB("src/server/databases/schedules.db")

	db.AutoMigrate(&StudentSchedule{})

	// TODO: Get only curr user elements
	result := db.Find(&list)
	if result.Error != nil {
		panic("failed to connect database")
	}
}

// Get returns the schedule
func Get() []StudentSchedule {
	return list
}

// Add creates and adds a class to the schedule
func Add(cid string) (string, error) {
	t := newStudentSchedule(cid)
	if cid == "" {
		return "", errors.New("message cannot be empty")
	}
	mtx.Lock()
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.ID, nil
}

// Delete removes and deletes a class from the schedule
func Delete(id string) error {
	location, err := findStudentScheduleLocation(id)
	if err != nil {
		return err
	}
	// TODO: Get UID
	db.Where("ID = ?", list[location].ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newStudentSchedule is a helper function to Add
func newStudentSchedule(cid string) StudentSchedule {
	// TODO: Get UID
	return StudentSchedule{
		ID:       xid.New().String(),
		User_ID:  "1005",
		Class_ID: cid,
	}
}

// findStudentScheduleLocation is a helper function to Delete
func findStudentScheduleLocation(id string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if t.ID == id {
			return i, nil
		}
	}
	return 0, errors.New("could not find schedule based on id")
}

// removeElementByLocation is a helper function to Delete
func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}
