package handlers

import (
	l "auth/internal/logger"
	"auth/internal/schemes"
	"auth/internal/tools"
	e "auth/internal/transport/http/error"
	"auth/internal/transport/http/response"
	sw "auth/internal/transport/http/swagger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var _ = sw.SwaggerAccessResponse{}

// @Summary Получить пару access, refresh токенов
// @tags Auth
// @Accept  json
// @Produce json
// @Param   guid          path int   true "Идентификатор пользователя"
// @Success 200 {object} swagger.SwaggerAccessResponse
// @Failure 400 {object} swagger.SwaggerNewError "guid должно быть числом > 0"
// @Failure 400 {object} swagger.SwaggerNewError "Ошибка создания токена"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга IP"
// @Failure 500 {object} swagger.SwaggerNewError "Ошибка парсинга UserAgent"
// @Router  /access/{guid} [post]
func (ah *authHandler) Access(w http.ResponseWriter, r *http.Request) {
	guid, err := strconv.Atoi(chi.URLParam(r, "guid"))
	if err != nil || guid <= 0 {
		response.NewResponse(
			e.NewError("guid должно быть числом > 0"),
			http.StatusBadRequest,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
		return
	}
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
	ac := schemes.AccessCreate{UserID: uint(guid), Ip: ip, UserAgent: useragent}

	ar, err := ah.service.Access(&ac)
	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка создания токена"),
			http.StatusBadRequest,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
		return
	}
	response.NewResponse(
		ar,
		http.StatusOK,
		w,
	)
}
