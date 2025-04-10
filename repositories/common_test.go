package repositories_test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
