package authdb

import "auth/internal/schemes"

func (ad *authDB) Create(data *schemes.AccessCreate) (string, error) {
	user, err := ad.userdb.GetUserById(data.UserID)
	if err != nil {
		return "", err
	}

	data.TokenVersion = user.TokenVersion

	err = ad.db.Create(data).Error
	if err != nil {
		return "", err
	}
	return data.Refresh, nil
}
