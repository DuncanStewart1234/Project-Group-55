// This package implements and maintains a list of users for the application
package user

import (
	"errors"
	"sync"
	"math/rand"
	"time"
	"strconv"

	"gorm.io/gorm"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
)

type UserType int
const (
	Student UserType = 0
	Admin   UserType = 1
)

var (
	list []User
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
	r *rand.Rand
)

// User is the struct used in this package to contain info about the User of this application
type User struct {
	gorm.Model
	ID         	uint
	User_ID    	int   		`json:"uid" gorm:"primaryKey"`
	First_Name 	string		`json:"first"`
	Last_Name  	string		`json:"last"`
	Email 	   	string		`json:"email"`
	User_Name  	string		`json:"uname"`
	Password	[]byte		`json:"hash"`
	Type       	UserType	`json:"utype"`
}

// init is a constructor that calls initialiseList
func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	once.Do(initialiseList)
}

// initialiseList creates the list of Users in this package and calls initDatabase
func initialiseList() {
	// list = []User{}
	initDatabase()
}

// initDatabase initialises the database for this package
func initDatabase() {
	db = utils.GetDB("src/server/databases/users.db")

	db.AutoMigrate(&User{})

	// // TODO: Return only curr user
	// result := db.Find(&list)
	// if result.Error != nil {
	// 	panic("failed to connect database")
	// }
}

// Get returns the list of Users in this package when called
// func Get() []User {
// 	return list
// }

// Add creates and stores a new User in the list and database
func Add(fname string, lname string, uname string, email string, pass string) (int, error) {
	mtx.Lock()
	if fname == "" || lname == "" {
		return 0, errors.New("name cannot be empty")
	}
	// TODO: Check if email/username already in database
	t := newUser(fname, lname, uname, email, pass)
	// list = append(list, t)
	db.Create(&t)
	mtx.Unlock()
	return t.User_ID, nil
}

//TODO: Add Edit Function

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
func newUser(fname string, lname string, uname string, email string, pass string) User {
	return User{
		User_ID:    r.Intn(89999999) + 10000000,
		First_Name: fname,
		Last_Name:  lname,
		Email: email,
		User_Name: uname,
		Password: utils.HashPasswrd(pass),
		Type: Student,
	}
}

// findUserLocation is a helper function to Delete
func findUserLocation(uid string) (int, error) {
	mtx.RLock()
	defer mtx.RUnlock()
	for i, t := range list {
		if strconv.Itoa(t.User_ID) == uid {
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

// func editUserByLocation(location int, fname string, lname string) {
// 	mtx.Lock()
// 	list[location].First_Name = fname
// 	list[location].Last_Name = lname
// 	mtx.Unlock()
// }