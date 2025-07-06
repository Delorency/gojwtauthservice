package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (as *authService) Access(data *schemes.AccessCreate) (*schemes.AccessResponse, error) {
	obj, f, err := as.repo.AuthorizedUserAgent(data.UserID, data.Ip)

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
		data.Refresh = tools.GetRefreshToken()

		err = as.repo.Create(data)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken(as.cfg, data.Jti, data.Ip, user.Email)
		if err != nil {
			return nil, err
		}
		return &schemes.AccessResponse{
			AccessToken:  access,
			RefreshToken: data.Refresh,
		}, nil
	}

	if obj.TokenVersion != obj.User.TokenVersion || time.Now().After(obj.ExpiredAt) {
		obj.TokenVersion = obj.User.TokenVersion
		obj.Jti = uuid.NewString()
		obj.ExpiredAt = time.Now().Add(as.cfg.Rtl)
		obj.Refresh = tools.GetRefreshToken()

		err = as.repo.Update(obj)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken(as.cfg, data.Jti, data.Ip, obj.User.Email)
		if err != nil {
			return nil, err
		}

		return &schemes.AccessResponse{
			AccessToken:  access,
			RefreshToken: obj.Refresh,
		}, nil
	}

	return nil, fmt.Errorf("")
}
