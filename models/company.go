package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	Name         string `json:"name" gorm:"unique;not null"`
	Description  string `json:"description"`
	Website      string `json:"website"`
	PrimaryEmail string `json:"email" gorm:"unique;not null"`

	Users []User `json:"users"` // Beziehung zu User
}
