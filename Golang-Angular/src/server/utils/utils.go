// This package contains a utility function for the sqlite databases
package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDB returns the specified database for use in other packages
func GetDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CheckIfEmptyOrTooLong(msg string) (error) {
	if msg == "" || len(msg) > 256 {
		return errors.New("string cannot be empty nor longer than 256 chars")
	}

	return nil
}

func HashPasswrd(password string) []byte {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to hash password")
	}

	return hashedPass
}