// This package contains a utility function for the sqlite databases
package utils

import (
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
