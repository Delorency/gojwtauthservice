package authdb

import (
	"auth/internal/models"

	"gorm.io/gorm"
)

func (ad *authDB) Delete(id uint) error {
	obj := models.RefreshToken{Model: gorm.Model{ID: id}}

	if err := ad.db.First(&obj).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return ad.db.Unscoped().Where("id = ?", id).Delete(&models.RefreshToken{}).Error
}
