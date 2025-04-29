package handlers

import (
	l "auth/internal/logger"
	"auth/internal/schemes"
	"auth/internal/tools"
	e "auth/internal/transport/http/error"
	"auth/internal/transport/http/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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
			e.NewError("Ошибка парсинга IP"),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}
	ac := schemes.AccessCreate{UserID: uint(guid), UserAgent: r.UserAgent(), Ip: ip}

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
