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

// @Summary Получение guid текущего пользователя
// @Security BearerAuth
// @tags Auth
// @Accept  json
// @Produce json
// @Success 200 {object} swagger.SwaggerMeResponse
// @Failure 401 {object} swagger.SwaggerNewError "Валидный токен не найден"
// @Failure 400 {object} swagger.SwaggerNewError "Неверный токен"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга IP"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга UserAgent"
// @Router  /me [get]
func (ah *authHandler) Me(w http.ResponseWriter, r *http.Request) {
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
	var req schemes.MeRequest

	req.Access = token

	ip, err := tools.GetIp(r)
	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка парсинга IP"),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}
	useragent, err := tools.GetUserAgent(r)
	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка парсинга UserAgent"),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}

	req.Ip = ip
	req.UserAgent = useragent

	id, err := ah.service.Me(&req)
	if err != nil {
		response.NewResponse(
			e.NewError("Неверный токен"),
			http.StatusBadRequest,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
		return
	}

	res := schemes.MeResponse{Guid: id}

	response.NewResponse(
		res,
		http.StatusOK,
		w,
	)
	ah.logger.Println(l.GetLogEntry(r, http.StatusOK, []byte{}))
}
