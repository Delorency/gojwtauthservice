package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (as *authService) Access(data *schemes.AccessCreate) (*schemes.AccessResponse, error) {
	obj, f, err := as.repo.GetByUserIDIPUserAgent(data.UserID, data.Ip, data.UserAgent)

	if err != nil {
		return nil, err
	}

	if !f {
		user, err := as.userdb.GetUserById(data.UserID)
		if err != nil {
			return nil, err
		}
		data.TokenVersion = user.TokenVersion

		data.Jti = uuid.NewString()
		data.ExpiredAt = time.Now().Add(as.cfg.Rtl)

		refreshtoken := tools.GetRefreshToken()

		bcrypthash, err := tools.GetBcryptHash(refreshtoken)
		if err != nil {
			return nil, err
		}

		data.Refresh = bcrypthash

		err = as.repo.Create(data)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken(as.cfg, data.UserID, data.Jti, data.Ip, data.UserAgent, user.Email)
		if err != nil {
			return nil, err
		}
		return &schemes.AccessResponse{
			AccessToken:  access,
			RefreshToken: refreshtoken,
		}, nil
	}

	if obj.TokenVersion != obj.User.TokenVersion || time.Now().After(obj.ExpiredAt) {
		obj.TokenVersion = obj.User.TokenVersion
		obj.Jti = uuid.NewString()
		obj.ExpiredAt = time.Now().Add(as.cfg.Rtl)

		refreshtoken := tools.GetRefreshToken()

		bcrypthash, err := tools.GetBcryptHash(refreshtoken)
		if err != nil {
			return nil, err
		}

		obj.Refresh = bcrypthash

		err = as.repo.Update(obj)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken(as.cfg, data.UserID, data.Jti, data.Ip, data.UserAgent, obj.User.Email)
		if err != nil {
			return nil, err
		}

		return &schemes.AccessResponse{
			AccessToken:  access,
			RefreshToken: refreshtoken,
		}, nil
	}

	return nil, fmt.Errorf("")
}
