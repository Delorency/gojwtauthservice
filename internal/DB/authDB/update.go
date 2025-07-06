package authdb

import (
	"auth/internal/models"

	"gorm.io/gorm"
)

func (ad *authDB) Update(data *models.RefreshToken) error {
	obj := models.RefreshToken{Model: gorm.Model{ID: data.ID}}

	return ad.db.Model(&obj).Updates(data).Error
}
