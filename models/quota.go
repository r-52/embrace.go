package models

import "gorm.io/gorm"

type Quota struct {
	gorm.Model
	Name      string  `json:"name" gorm:"unique;not null"`
	CompanyID uint    `json:"-"`
	Company   Company `json:"company"`

	Count        int    `json:"count" gorm:"not null"`
	QuotaResetAt string `json:"quotaResetAt" gorm:"not null,default:'firstOfYear'"`
}

const QUOTA_RESET_FIRST_OF_YEAR = "firstOfYear"
const QUOTA_RESET_FIRST_OF_MONTH = "firstOfMonth"
const QUOTA_RESET_FIRST_OF_WEEK = "firstOfWeek"
