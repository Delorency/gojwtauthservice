package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Jti          string `gorm:"uniqueIndex,size:128"`
	Refresh      string `gorm:"size:128"`
	ExpiredAt    time.Time
	Ip           string `gorm:"size:64"`
	UserAgent    string `gorm:"size:128"`
	TokenVersion string `gorm:"size:128"`

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE"`
}

func (RefreshToken) TableName() string {
	return "refreshtoken"
}
