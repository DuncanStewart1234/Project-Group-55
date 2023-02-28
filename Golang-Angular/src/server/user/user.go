package schedule

import (
	"errors"
	"sync"
    "gorm.io/gorm"

    "github.com/rs/xid"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

var (
	list []User
	db *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

type User struct {
	gorm.Model
	User_ID		string `json:"cid"`
	First_Name  string `json:"first"`
	Last_Name   string `json:"last"`
}

func init() {
	once.Do(initialiseList)
}

func initialiseList() {
	list = []User{}
	initDatabase()
}

func initDatabase() {
	db = utils.GetDB("src/server/databases/users.db")

    db.AutoMigrate(&User{})

	result := db.Find(&list)
    if result.Error	!= nil {
        panic("failed to connect database")
    }
}

func Get() []User {
	return list
}

func Add(fname string, lname string) (string, error) {
	t := newUser(fname, lname)
	if fname == "" || lname == "" {
		return "", errors.New("name cannot be empty")
	}
	mtx.Lock()
	list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.User_ID, nil
}

func Delete(uid string) error {
	location, err := findUserLocation(uid)
	if err != nil {
		return err
	}
	db.Where("User_ID = ?", list[location].User_ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

func newUser(fname string, lname string) User {
	return User {
		User_ID: xid.New().String(),
		First_Name: fname,
		Last_Name: lname,
	}
}

func findUserLocation(uid string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if t.User_ID == uid {
			return i, nil
		}
	}
	return 0, errors.New("could not find user based on uid")
}

func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}