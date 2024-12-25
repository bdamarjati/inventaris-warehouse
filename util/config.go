package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
