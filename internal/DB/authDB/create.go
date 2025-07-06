package authdb

import (
	"auth/internal/models"
	"auth/internal/schemes"
)

func (ad *authDB) Create(data *schemes.AccessCreate) error {
	obj := &models.RefreshToken{
		Jti:          data.Jti,
		Refresh:      data.Refresh,
		ExpiredAt:    data.ExpiredAt,
		Ip:           data.Ip,
		TokenVersion: data.TokenVersion,
		UserID:       data.UserID,
	}
	return ad.db.Create(obj).Error
}
