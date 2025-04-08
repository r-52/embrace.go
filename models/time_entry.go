package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type TimeEntry struct {
	gorm.Model

	StartTime time.Time       `json:"startTime" gorm:"not null"`
	EndTime   sql.NullTime    `json:"endTime"`
	Duration  sql.NullFloat64 `json:"duration"`
	Note      string          `json:"note"`

	UserID uint `json:"-"`
	User   User `json:"user"`

	TimeEntryTypeID uint          `json:"-"`
	TimeEntryType   TimeEntryType `json:"timeEntryType"`
}

type TimeEntryType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`

	Color       string         `json:"color" gorm:"not null"`
	Icon        sql.NullString `json:"icon"`
	Description sql.NullString `json:"description"`

	CompanyID uint    `json:"-"`
	Company   Company `json:"company"`
}
