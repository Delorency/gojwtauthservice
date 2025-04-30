package authdb

import (
	"auth/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (ad *authDB) AuthorizedUserAgent(userid uint, useragent string, ip string) (*models.RefreshToken, bool, error) {
	var refreshdata models.RefreshToken
	result := ad.db.Where("user_agent = ? and user_id = ? and ip = ?", useragent, userid, ip).Preload("User").First(&refreshdata)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, nil

		}
		return nil, false, result.Error
	}
	return &refreshdata, true, nil
}
func (ad *authDB) AuthorizedUserToken(refresh string, useragent string) (*models.RefreshToken, error) {
	var refreshdata models.RefreshToken
	err := ad.db.Where("refresh = ? and user_agent = ?", refresh, useragent).Preload("User").First(&refreshdata).Error
	if err != nil {
		return nil, err
	}
	return &refreshdata, nil
}
