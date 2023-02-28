package schedule

import (
	"errors"
	"sync"
    "gorm.io/gorm"

    "github.com/rs/xid"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []StudentSchedule
	db *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

type StudentSchedule struct {
	gorm.Model
	ID       string `json:"id"`
	User_ID	 string `json:"uid"`
	Class_ID string `json:"cid"`
}

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []StudentSchedule{}
	initDatabase()
}

func initDatabase() {
	db = utils.GetDB("src/server/databases/schedules.db")

    db.AutoMigrate(&StudentSchedule{})

	// TODO: Get only curr user elements
	result := db.Find(&list)
    if result.Error	!= nil {
        panic("failed to connect database")
    }
}

func Get() []StudentSchedule {
	return list
}

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

func newStudentSchedule(cid string) StudentSchedule {
	// TODO: Get UID
	return StudentSchedule {
		ID:       xid.New().String(),
		User_ID: "1005",
		Class_ID:  cid,
	}
}

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

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}