package authdb

import (
	"auth/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (ad *authDB) GetByUserIDAndIP(userid uint, ip string) (*models.RefreshToken, bool, error) {
	var refreshdata models.RefreshToken
	result := ad.db.Where("user_id = ? and ip = ?", userid, ip).Preload("User").First(&refreshdata)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, nil

		}
		return nil, false, result.Error
	}
	return &refreshdata, true, nil
}
func (ad *authDB) GetByToken(bcryptrefresh string) (*models.RefreshToken, error) {
	var refreshdata models.RefreshToken
	err := ad.db.Where("refresh = ?", bcryptrefresh).Preload("User").First(&refreshdata).Error
	if err != nil {
		return nil, err
	}
	return &refreshdata, nil
}
