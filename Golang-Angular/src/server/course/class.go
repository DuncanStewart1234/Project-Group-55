// A package used to create a schedule of classes that a college student
// may use.
package course

import (
	"errors"
	"sync"
	"math/rand"
	"time"
	"strconv"

	"gorm.io/gorm"
	"encoding/json"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []Class
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
	r *rand.Rand
)

// Class is a struct used to contain info about a student's class
type Class struct {
	gorm.Model
	ID uint
	Class_ID int `json:"cid" gorm:"primaryKey"`
	Name string `json:"name"`
	Abbrv string `json:"abbrv"`
	Location Location `json:"loc" gorm:"serializer:json"`
	Schedule Schedule `json:"schedule" gorm:"serializer:json"`
}

type Location struct {
	Lat float64
	Long float64
}

type Period struct {
	S_Time string
	E_Time string
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
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
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
func Add(name string, abbrv string, loc string, scheduleBlock string) (int, error) {
	t := newClass(name, abbrv, getLocationFromJSON(loc), getScheduleFromJSON(scheduleBlock))
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
func newClass(name string, abbrv string, loc Location, s Schedule) Class {
	return Class{
		Class_ID: r.Intn(899999) + 100000,
		Name: name,
		Abbrv: abbrv,
		Location: loc,
		Schedule: s,
	}
}

func newPeriod(l [][]string) []Period {
	var periods []Period
	for _, t := range l {
		periods = append(periods, Period{
			S_Time: t[0],
			E_Time: t[1],
		})
	}

	return periods
}

func getLocationFromJSON(jsonLoc string) Location {
	var loc []float64
	json.Unmarshal([]byte(jsonLoc), &loc)
	return Location{
		Lat: loc[0],
		Long: loc[1],
	}
}

func getScheduleFromJSON(jsonSchedule string) Schedule {
	var data map[string][][]string
	err := json.Unmarshal([]byte(jsonSchedule), &data)
	if err != nil {
		panic(err)
	}

	s := Schedule {
		Mon: newPeriod(data["Mon"]),
		Tues: newPeriod(data["Tues"]),
		Wed: newPeriod(data["Wed"]),
		Thur: newPeriod(data["Thur"]),
		Fri: newPeriod(data["Fri"]),
		Sat: newPeriod(data["Sat"]),
		Sun: newPeriod(data["Sun"]),
	}

	return s
}

// findClassLocation is a helper function to Delete
func findClassLocation(cid string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if strconv.Itoa(t.Class_ID) == cid {
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
