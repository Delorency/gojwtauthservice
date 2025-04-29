package authdb

import (
	"auth/internal/models"
	"auth/internal/schemes"
	"errors"

	"gorm.io/gorm"
)

func (ad *authDB) Access(data *schemes.AccessCreate) (string, error) {
	var refreshdata models.RefreshToken
	result := ad.db.Where("useragent = ?", data.UserAgent).First(&refreshdata)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return ad.Create(data)

		}
		return "", result.Error
	}

	obj := models.RefreshToken{UserID: data.UserID}

	if err := ad.db.Model(&obj).Updates(data).Error; err != nil {
		return "", err
	}

	return data.Refresh, nil
}
