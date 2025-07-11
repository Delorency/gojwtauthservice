package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
