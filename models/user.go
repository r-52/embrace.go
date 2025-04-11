package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`

	CompanyID uint    `json:"-"`
	Company   Company `json:"company"`

	RoleID uint     `json:"-"`
	Role   UserRole `json:"role"`

	UserProfile   UserProfile `json:"userProfile"`
	UserProfileID uint        `json:"-"`
	TimeEntries   []TimeEntry `json:"timeEntries"`
}

type UserQuota struct {
	gorm.Model
	QuotaID uint  `json:"-"`
	Quota   Quota `json:"quota"`
	UserID  uint  `json:"-"`
	User    User  `json:"user"`
	Count   int   `json:"count" gorm:"not null"`
}

type UserProfile struct {
	gorm.Model

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
	Position  string `json:"position"`
	Location  string `json:"location"`

	Slug   string `json:"slug" gorm:"unique;not null"`
	Avatar string `json:"avatar"`
	Phone  string `json:"phone"`
}

type UserRole struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`

	InternalUsage int `json:"internalUsage" gorm:"default:0"`
}
