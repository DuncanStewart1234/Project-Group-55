// This package implements and maintains a list of users for the application
package user

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils"
	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/utils/token"
)

var (
	curr_user User
	db   *gorm.DB
	mtx  sync.RWMutex
	once sync.Once
	r *rand.Rand
	// currToken string
)

// User is the struct used in this package to contain info about the User of this application
type User struct {
	// gorm.Model
	// ID         	uint
	User_ID    	int   		`json:"uid" gorm:"primaryKey"`
	First_Name 	string		`json:"first" gorm:"not null"`
	Last_Name  	string		`json:"last" gorm:"not null"`
	Email 	   	string		`json:"email" gorm:"not null;unique"`
	User_Name  	string		`json:"uname" gorm:"not null;unique"`
	Password	[]byte		`json:"password" gorm:"not null"`
}

// init is a constructor that calls initialiseList
func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	once.Do(initialiseList)
}

// initialiseList creates the list of Users in this package and calls initDatabase
func initialiseList() {
	initDatabase()
}

// initDatabase initialises the database for this package
func initDatabase() {
	envErr := godotenv.Load("src/server/.env")
	if envErr != nil {
		panic(envErr)
	}

	db = utils.GetDB("src/server/databases/users.db")

	db.AutoMigrate(&User{})
}

func Get(uname string) (User) {
	returnUser := curr_user
	if uname != returnUser.User_Name {
		panic("not correct user")
	}

	returnUser.Password = nil
	return returnUser
}

func GetAll() ([]User) {
	var list []User
	db.Find(&list)

	return list
}

// Add creates and stores a new User in the list and database
func Add(fname string, lname string, uname string, email string, pass string) (int, error) {
	mtx.Lock()
	if fname == "" || lname == "" {
		return 0, errors.New("name cannot be empty")
	}

	t := newUser(fname, lname, uname, email, pass)
	result := db.Create(&t)
	if result.Error != nil {
		return -1, result.Error
	}
	mtx.Unlock()
	return t.User_ID, nil
}

func Login(uname string, passwd string) (string, error) {
	result := db.Where("User_Name = ?", uname).Take(&curr_user)
	if result.Error != nil {
		return "", result.Error
	}
	
	hashErr := utils.CheckHashedPasswrd(passwd, curr_user.Password)
	if hashErr != nil {
		return "", hashErr
	}
	
	token, tokenErr := token.GenerateToken(curr_user.User_ID)
	if tokenErr != nil {
		return "", tokenErr
	}
	
	// currToken = token
	return token, nil
}

//TODO: Add Edit Function
func Edit(uid string, fname string, lname string, uname string, email string, new_pass string, old_pass string) error {
	result := db.Where("User_ID = ?", uid).Take(&curr_user)
	if result.Error != nil {
		return result.Error
	}
	
	hashErr := utils.CheckHashedPasswrd(old_pass, curr_user.Password)
	if hashErr != nil {
		return hashErr
	}

	mtx.Lock()
	if fname != "" {
		curr_user.First_Name = fname
	}
	if lname != "" {
		curr_user.Last_Name = lname
	}
	if email != "" {
		curr_user.Email = email
	}
	if uname != "" {
		curr_user.User_Name = utils.Fix_Username(uname)
	}
	if new_pass != "" {
		curr_user.Password = utils.HashPasswrd(new_pass)
	}
	mtx.Unlock()
	db.Save(&curr_user)

	return nil
}

// Delete removes a User from the list and deletes them from the database
func Delete(uid string) error {
	uidValue, err := strconv.Atoi(uid)
	if err != nil || uidValue != curr_user.User_ID {
		return errors.New("cannot delete this account")
	}
	db.Where("User_ID = ?", curr_user.User_ID).Unscoped().Delete(&curr_user)
	return nil
}

func Logout() {
	var nullUser User
	curr_user = nullUser
}

// newUser is a helper function to Add
func newUser(fname string, lname string, uname string, email string, pass string) User {
	return User{
		User_ID:	r.Intn(89999999) + 10000000,
		First_Name:	fname,
		Last_Name:	lname,
		Email:		email,
		User_Name:	utils.Fix_Username(uname),
		Password:	utils.HashPasswrd(pass),
	}
}

func GetUID() (int) {
	if curr_user.User_ID == 0 {
		panic("not logged in")
	}
	return curr_user.User_ID
}