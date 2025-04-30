package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"size:128"`
	TokenVersion string `gorm:"size:128;default:'v1'"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
