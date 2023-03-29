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
	curr_uid int
)

// StudentSchedule is the struct used to create a class schedule
type StudentSchedule struct {
	gorm.Model
	ID       string `json:"id"`
	User_ID  int `json:"uid"`
	Class_ID string `json:"cid"`
}

// init is a constructor, calls initialiseList
func init() {
	once.Do(initialiseList)
}

// initialiseList creates the schedule array and calls initDatabase
func initialiseList() {
	list = []StudentSchedule{}
	// TODO: Get curr user uid
	curr_uid = 1005
	initDatabase()
}

// initDatabase initialises the database
func initDatabase() {
	db = utils.GetDB("src/server/databases/schedules.db")

	db.AutoMigrate(&StudentSchedule{})

	result := db.Where("User_ID = ?", curr_uid).Find(&list)
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
	// TODO: Check if cid in ClassDB
	mtx.Lock()
	t := newStudentSchedule(cid)
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
	db.Where("ID = ?", list[location].ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newStudentSchedule is a helper function to Add
func newStudentSchedule(cid string) StudentSchedule {
	return StudentSchedule{
		ID:       xid.New().String(),
		User_ID:  curr_uid,
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
