// This package implements and maintains a list of users for the application
package schedule

import (
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/rs/xid"
)

var (
	list []User
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
)

// User is the struct used in this package to contain info about the User of this application
type User struct {
	gorm.Model
	User_ID    string `json:"cid"`
	First_Name string `json:"first"`
	Last_Name  string `json:"last"`
}

// init is a constructor that calls initialiseList
func init() {
	once.Do(initialiseList)
}

// initialiseList creates the list of Users in this package and calls initDatabase
func initialiseList() {
	list = []User{}
	initDatabase()
}

// initDatabase initialises the database for this package
func initDatabase() {
	db = utils.GetDB("src/server/databases/users.db")

	db.AutoMigrate(&User{})

	result := db.Find(&list)
	if result.Error != nil {
		panic("failed to connect database")
	}
}

// Get returns the list of Users in this package when called
func Get() []User {
	return list
}

// Add creates and stores a new User in the list and database
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

// Delete removes a User from the list and deletes them from the database
func Delete(uid string) error {
	location, err := findUserLocation(uid)
	if err != nil {
		return err
	}
	db.Where("User_ID = ?", list[location].User_ID).Delete(&list[location])
	removeElementByLocation(location)
	return nil
}

// newUser is a helper function to Add
func newUser(fname string, lname string) User {
	return User{
		User_ID:    xid.New().String(),
		First_Name: fname,
		Last_Name:  lname,
	}
}

// findUserLocation is a helper function to Delete
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

// removeElementByLocation is a helper function to Delete
func removeElementByLocation(i int) {
	mtx.Lock()
	list = append(list[:i], list[i+1:]...)
	mtx.Unlock()
}
