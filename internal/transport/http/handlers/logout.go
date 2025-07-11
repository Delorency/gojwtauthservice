package handlers

import (
	l "auth/internal/logger"
	"auth/internal/schemes"
	"auth/internal/tools"
	e "auth/internal/transport/http/error"
	"auth/internal/transport/http/response"
	sw "auth/internal/transport/http/swagger"
	"net/http"
)

var _ = sw.SwaggerAccessResponse{}

// @Summary Осуществить logout по access токену
// @Security BearerAuth
// @tags Auth
// @Accept  json
// @Produce json
// @Success 204 "Успешно"
// @Failure 401 {object} swagger.SwaggerNewError "Валидный токен не найден"
// @Failure 400 {object} swagger.SwaggerNewError "Ошибка выхода"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга IP"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга UserAgent"
// @Router  /logout [post]
func (ah *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := tools.GetTokenFromHeader(r)
	if err != nil {
		response.NewResponse(
			e.NewError("Валидный токен не найден"),
			http.StatusUnauthorized,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusUnauthorized, []byte{}))
		return
	}
	var req schemes.LogoutRequest

	req.Access = token

	ip, err := tools.GetIp(r)
	if err != nil {
		response.NewResponse(
			e.NewError(err.Error()),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}
	useragent, err := tools.GetUserAgent(r)
	if err != nil {
		response.NewResponse(
			e.NewError(err.Error()),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}

	req.Ip = ip
	req.UserAgent = useragent

	err = ah.service.Logout(&req)
	if err != nil {
		response.NewResponse(
			e.NewError(err.Error()),
			http.StatusBadRequest,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
		return
	}

	response.NewResponse(
		[]byte{},
		http.StatusNoContent,
		w,
	)
	ah.logger.Println(l.GetLogEntry(r, http.StatusOK, []byte{}))
}
