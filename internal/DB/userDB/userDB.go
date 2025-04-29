package userdb

import (
	"auth/internal/models"

	"gorm.io/gorm"
)

type UserDBI interface {
	GetUserById(uint) (*models.User, error)
}

type userDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) UserDBI {
	return &userDB{db}
}
