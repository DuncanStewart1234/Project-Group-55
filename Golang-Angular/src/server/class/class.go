package schedule

import (
	"errors"
	"sync"
    "gorm.io/gorm"

    "github.com/rs/xid"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []Class
	db *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

type Class struct {
	gorm.Model
	// ID       string `json:"id"`
	Class_ID string `json:"cid"`
	Location string `json:"loc"`
	// Add Schedule
}

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []Class{}
	initDatabase()
}

func initDatabase() {
	db = utils.GetDB("src/server/databases/classes.db")

    db.AutoMigrate(&Class{})

	result := db.Find(&list)
    if result.Error	!= nil {
        panic("failed to connect database")
    }
}

func Get() []Class {
	return list
}

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

func Delete(cid string) error {
	location, err := findClassLocation(cid)
	if err != nil {
		return err
	}
	db.Where("Class_ID = ?", list[location].Class_ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

func newClass(loc string) Class {
	return Class {
		Class_ID: xid.New().String(),
		Location: loc,
	}
}

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

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}