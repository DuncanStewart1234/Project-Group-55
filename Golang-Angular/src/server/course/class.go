// A package used to create a schedule of classes that a college student
// may use.
package course

import (
	"errors"
	"sync"

	"gorm.io/gorm"
	"encoding/json"
	"time"

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
	Class_ID string `json:"cid"`
	Name string `json:"name"`
	Abbrv string `json:"abbrv"`
	Location []byte `json:"loc"`
	Schedule []byte `json:"schedule"`
}

type Location struct {
	Lat float64
	Long float64
}

type Period struct {
	// TODO: TIME + DURATION OR STARTTIME + ENDTIME
	S_Time time.Time
	E_Time time.Time
}

type Schedule struct {
	Mon []Period `json:"Mon"`
    Tues []Period `json:"Tues"`
    Wed []Period `json:"Wed"`
    Thur []Period `json:"Thur"`
    Fri []Period `json:"Fri"`
    Sat []Period `json:"Sat"`
    Sun []Period `json:"Sun"`
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
	db = utils.GetDB("src/server/databases/classes.db")

	db.AutoMigrate(&Class{})

	result := db.Find(&list)
	// TODO: Convert Schedule and location
	if result.Error != nil {
		panic("failed to connect database")
	}
}

// Get returns the list of classes in the schedule
func Get() []Class {
	return list
}

// Add creates and adds a class element to the list
// TODO: Define scheduleBlock using frontend API
func Add(name string, abbrv string, lat float64, long float64, scheduleBlock string) (string, error) {
	t := newClass(name, abbrv, lat, long, []byte(scheduleBlock))
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
func newClass(name string, abbrv string, lat float64, long float64, s []byte) Class {
	return Class{
		Class_ID: xid.New().String(),
		Name: name,
		Abbrv: abbrv,
		Location: newLocation(lat, long),
		Schedule: s,
	}
}

func newLocation(lat float64, long float64) []byte {
	loc := &Location{
		Lat: lat,
		Long: long,
	}
	b, _ := json.Marshal(loc)
	return b
}

// func newSchedule(blocks string) []byte {
// 	// periods := []Period
// 	// TODO: Complete function
// 	var m Message
// 	l := Schedule{
// 	}
// 	b, _ := json.Marshal(l)
// 	return b
// }

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
