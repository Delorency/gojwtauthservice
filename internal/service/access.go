package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (as *authService) Access(data *schemes.AccessCreate) (*schemes.AccessResponse, error) {
	obj, f, err := as.repo.AuthorizedUserAgent(data.UserID, data.UserAgent)

	if err != nil {
		return nil, err
	}

	// если устройство не найдено
	if !f {
		user, err := as.userdb.GetUserById(data.UserID)
		if err != nil {
			return nil, err
		}
		data.TokenVersion = user.TokenVersion

		data.Jti = uuid.NewString()
		data.ExpiredAt = time.Now().Add(as.cfg.Rtl)
		data.Refresh = tools.EncodeToBase64()

		refresh, err := as.repo.Create(data)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken()
		if err != nil {
			return nil, err
		}
		return &schemes.AccessResponse{
			AccessToken:  access,
			RefreshToken: refresh,
		}, nil
	}
	if obj.TokenVersion != obj.User.TokenVersion || tools.CheckExpire(obj.ExpiredAt) {
		obj.TokenVersion = obj.User.TokenVersion
		obj.Jti = uuid.NewString()
		obj.ExpiredAt = time.Now().Add(as.cfg.Rtl)
		obj.Refresh = tools.EncodeToBase64()

		err := as.repo.Update(obj)
		if err != nil {
			return nil, err
		}

		access, err := tools.GetJWTToken()
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
