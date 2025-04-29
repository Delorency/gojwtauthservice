package userdb

import (
	"auth/internal/models"

	"gorm.io/gorm"
)

func (ud *userDB) GetUserById(id uint) (*models.User, error) {
	obj := models.User{Model: gorm.Model{ID: id}}

	err := ud.db.First(&obj).Error

	if err != nil {
		return nil, err
	}

	return &obj, err
}
