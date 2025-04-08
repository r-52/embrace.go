package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
	Position  string `json:"position"`
	Location  string `json:"location"`

	CompanyID uint    `json:"-"`
	Company   Company `json:"company"`

	Slug   string `json:"slug" gorm:"unique;not null"`
	Avatar string `json:"avatar"`
	Phone  string `json:"phone"`

	RoleID uint     `json:"-"`
	Role   UserRole `json:"role"`
}

type UserRole struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
