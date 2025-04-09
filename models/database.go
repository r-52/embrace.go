package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	if dbConnection == "" {
		dbConnection = "test.db"
	}
	db, err := gorm.Open(sqlite.Open(dbConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Company{}, &User{}, &UserRole{}, &TimeEntry{}, &TimeEntryType{}, &UserProfile{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
