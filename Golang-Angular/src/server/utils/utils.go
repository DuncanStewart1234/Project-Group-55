package utils

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func GetDB(path string) (*gorm.DB) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	return db
}