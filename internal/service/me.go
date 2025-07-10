package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"strings"
)

func (as *authService) Me(data *schemes.MeRequest) (uint, error) {
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return 0, fmt.Errorf("")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])

	if !f {
		return 0, fmt.Errorf("Parse token payload error")
	}

	if payload.UserAgent != data.UserAgent && payload.Ip != data.Ip {
		return 0, fmt.Errorf("Incorrect token")
	}

	return payload.Id, nil
}
