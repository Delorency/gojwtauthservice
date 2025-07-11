package authdb

import (
	"auth/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (ad *authDB) GetByUserIDIPUserAgent(userid uint, ip string, useragent string) (*models.RefreshToken, bool, error) {
	var refreshdata models.RefreshToken
	result := ad.db.Where("user_id = ? and ip = ? and user_agent = ?", userid, ip, useragent).Preload("User").First(&refreshdata)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, nil

		}
		return nil, false, result.Error
	}
	return &refreshdata, true, nil
}
