package handlers

import (
	l "auth/internal/logger"
	"auth/internal/schemes"
	"auth/internal/tools"
	e "auth/internal/transport/http/error"
	"auth/internal/transport/http/response"
	"encoding/json"
	"io"
	"net/http"

	sw "auth/internal/transport/http/swagger"
	v "auth/internal/validator"

	"github.com/go-playground/validator"
)

var _ = sw.SwaggerAccessResponse{}

// @Summary Получить новый access токен
// @Security BearerAuth
// @tags Auth
// @Accept  json
// @Produce json
// @Param   person body swagger.SwaggerRefreshRequest true "Refresh токен"
// @Success 201 {object} swagger.SwaggerAccessResponse
// @Failure 401 {object} swagger.SwaggerNewError "Валидный токен не найден"
// @Failure 400 {object} swagger.SwaggerValidateData "Необходимые поля не заполнены"
// @Failure 400 {object} swagger.SwaggerNewError "Ошибка обновления токена"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга IP"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга тела запроса"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга UserAgent"
// @Router  /refresh [post]
func (ah *authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
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
	bodyBytes, _ := io.ReadAll(r.Body)
	var req schemes.RefreshRequest

	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		response.NewResponse(
			e.NewError("Необходимые поля не заполнены"),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}
	if err := validator.New().Struct(req); err != nil {
		data, f := v.HandleValidationErrors(w, err)
		if f {
			response.NewResponse(
				data,
				http.StatusBadRequest,
				w,
			)
			ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, bodyBytes))
			return
		}
	}
	ip, err := tools.GetIp(r)
	if err != nil {
		response.NewResponse(
			e.NewError(err.Error()),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
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

	req.Access = token

	ar, err := ah.service.Refresh(&req)
	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка обновления токена"),
			http.StatusBadRequest,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, bodyBytes))
		return
	}
	response.NewResponse(
		ar,
		http.StatusOK,
		w,
	)
	ah.logger.Println(l.GetLogEntry(r, http.StatusOK, bodyBytes))
}
